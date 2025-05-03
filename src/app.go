package main

import (
	"scrapher/src/config"
	"scrapher/src/global"
	"scrapher/src/middleware"

	"scrapher/src/modules"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

var service = "Scrapher Service"

// Initializes the Fiber application with middleware, routes, and database connection.
func bootstrapApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:           service,
		EnablePrintRoutes: true,
		ErrorHandler:      middleware.ErrorHandler,
		BodyLimit:         50 * 1024, // 50 KB, for now we don't need much since we are not sending large payloads
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e any) {
			log.Error(e, zap.Stack("stacktrace"))
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Env.FrontendBaseUrl,
		AllowCredentials: true,
	}))

	app.Use(helmet.New())

	app.Use(requestid.New(requestid.Config{
		Header:     global.HdrXCorrelationID,
		ContextKey: global.CtxCorrelationID,
	}))

	app.Use(middleware.Zapped)

	app.Use(middleware.ResponseInterceptors...)

	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(global.Response[*interface{}]{
				Message: "Too many requests, please try again later",
			})
		},
	}))

	app.Use(middleware.HealthCheck(middleware.HealthCheckOptions{
		Service:        &service,
		CheckFunctions: map[string]func() bool{},
	}))

	app.Get("/metrics", monitor.New())

	app.Mount("/api", modules.New())

	return app
}
