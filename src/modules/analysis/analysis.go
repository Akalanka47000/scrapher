package analysis

import (
	v1 "scrapher/src/modules/analysis/api/v1"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	analysis := fiber.New()
	analysis.Mount("/v1", v1.New())
	return analysis
}
