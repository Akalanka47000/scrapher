package v1

import (
	m "scrapher/src/middleware"
	"scrapher/src/modules/analysis/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	v1 := fiber.New()

	v1.Post("/", m.Zelebrate[dto.PerformAnalysisRequest](m.ZelebrateSegmentBody), PerformAnalysis)

	return v1
}
