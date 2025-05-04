package middleware

import (
	"scrapher/src/config"
	"scrapher/src/global"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// Protects an API route by checking if the request contains a valid service request key.
func Sentinel(ctx *fiber.Ctx) error {
	if config.Env.ServiceRequestKey == "" {
		return ctx.Next()
	}
	if lo.CoalesceOrEmpty(ctx.Get(global.HdrXServiceRequestKey), ctx.Query("token")) !=
		config.Env.ServiceRequestKey {
		panic(fiber.NewError(fiber.StatusForbidden, "You are not permitted to access this resource"))
	}
	return ctx.Next()
}
