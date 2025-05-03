package modules

import (
	"scrapher/src/modules/analysis"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	modules := fiber.New()

	modules.Mount("/analysis", analysis.New())

	return modules
}
