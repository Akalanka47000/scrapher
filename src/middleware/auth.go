package middleware

import (
	"scrapher/src/config"
	"scrapher/src/global"

	"github.com/gofiber/fiber/v2"
)

func Sentinel(ctx *fiber.Ctx) error {
	if config.Env.ServiceRequestKey != "" {
		if ctx.Get(global.HdrXServiceRequestKey) != config.Env.ServiceRequestKey {
			panic(fiber.NewError(fiber.StatusForbidden, "You are not permitted to access this resource"))
		}
	}
	return ctx.Next()
}
