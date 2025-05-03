package middleware

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/samber/lo"
)

var prefix = "system"

type HealthCheckOptions struct {
	// Service name to be displayed in the health check response
	Service *string
	// Auth middleware to be execute before health checks. Use this to protect the health check endpoints from spammers
	AuthMiddleware *fiber.Handler
	// CheckFunctions is a map of functions to be executed for health checks
	CheckFunctions map[string]func() bool
	// Custom routes for health checks
	Routes struct {
		Health    string
		Readiness string
		Liveness  string
	}
}

func probeHandler(c *fiber.Ctx, checkFunctions map[string]func() bool) error {
	status := fiber.StatusOK
	var wg sync.WaitGroup
	results := make(map[string]bool)
	for k, check := range checkFunctions {
		wg.Add(1)
		go func(check func() bool) {
			defer wg.Done()
			key := fmt.Sprintf("%s_check", k)
			results[key] = check()
			if !results[key] {
				status = fiber.StatusServiceUnavailable
			}
		}(check)
	}
	wg.Wait()

	return c.Status(status).JSON(results)
}

func HealthCheck(opts HealthCheckOptions) fiber.Handler {
	opts.Routes.Health = lo.CoalesceOrEmpty(opts.Routes.Health, fmt.Sprintf("/%s/health", prefix))
	opts.Routes.Readiness = lo.CoalesceOrEmpty(opts.Routes.Readiness, fmt.Sprintf("/%s/readiness", prefix))
	opts.Routes.Liveness = lo.CoalesceOrEmpty(opts.Routes.Liveness, fmt.Sprintf("/%s/liveness", prefix))
	if len(opts.CheckFunctions) == 0 {
		opts.CheckFunctions = map[string]func() bool{
			"app": func() bool { return true },
		}
	}
	return func(c *fiber.Ctx) error {
		if c.Method() != fiber.MethodGet {
			return c.Next()
		}

		prefixCount := len(utils.TrimRight(c.Route().Path, '/'))

		if opts.AuthMiddleware != nil {
			if err := (*opts.AuthMiddleware)(c); err != nil {
				return err
			}
		}

		if len(c.Path()) >= prefixCount {
			checkPath := c.Path()[prefixCount:]
			checkPathTrimmed := checkPath
			if !c.App().Config().StrictRouting {
				checkPathTrimmed = utils.TrimRight(checkPath, '/')
			}
			switch {
			case checkPath == opts.Routes.Health || checkPathTrimmed == opts.Routes.Health:
				message := "OK"
				if opts.Service != nil {
					message = fmt.Sprintf("%s - %s", *opts.Service, message)
				}
				return c.Status(fiber.StatusOK).SendString(message)
			case checkPath == opts.Routes.Readiness || checkPathTrimmed == opts.Routes.Readiness:
				return probeHandler(c, opts.CheckFunctions)
			case checkPath == opts.Routes.Liveness || checkPathTrimmed == opts.Routes.Liveness:
				return probeHandler(c, opts.CheckFunctions)
			}
		}

		return c.Next()
	}
}
