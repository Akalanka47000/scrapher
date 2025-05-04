package v1

import (
	m "scrapher/src/middleware"
	"scrapher/src/modules/analysis/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	v1 := fiber.New()

	v1.Get("/webpage",
		m.Zelebrate[dto.AnalyseWebpageRequest](m.ZelebrateSegmentQuery),
		m.CacheSuccess, AnalyseWebpage)

	return v1
}
