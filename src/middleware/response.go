package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"os"
	"scrapher/src/global"
)

// Intercepts the response and populates it with additional fields.
func headerInterceptor(c *fiber.Ctx) error {
	c.Append(global.HdrXHostname, lo.Ok(os.Hostname()))
	return c.Next()
}

var ResponseInterceptors = []any{headerInterceptor}
