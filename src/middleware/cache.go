package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

// Caches successful responses for 30 minutes. Please refer the fiber cache documentation
// for more details: https://docs.gofiber.io/api/middleware/cache
var CacheSuccess = cache.New(cache.Config{
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.OriginalURL()
	},
	Expiration: time.Minute * 30,
})
