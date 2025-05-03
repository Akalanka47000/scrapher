package analysis

import (
	"github.com/akalanka47000/go-modkit/routing"
	v1 "scrapher/src/modules/analysis/api/v1"

	"github.com/gofiber/fiber/v2"
)

var versioned = routing.VersionablePrefix("analysis")

func New() *fiber.App {
	analysis := fiber.New()
	analysis.Mount(versioned(1), v1.New())
	return analysis
}
