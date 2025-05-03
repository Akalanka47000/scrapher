package middleware

import (
	"os"
	"scrapher/src/global"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// Intercepts the response and populates it with additional fields.
func responseHeaderInjector(c *fiber.Ctx) error {
	c.Append(global.HdrXHostname, lo.Ok(os.Hostname())) // Useful for debugging
	return c.Next()
}

var Injectors = []any{responseHeaderInjector}
