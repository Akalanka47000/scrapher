package v1

import (
	"scrapher/src/global"
	"scrapher/src/modules/analysis/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func PerformAnalysis(c *fiber.Ctx) error {
	payload := new(dto.PerformAnalysisRequest)
	c.BodyParser(payload)
	performAnalysis(c, payload.TargetURL)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.PerformAnalysisResponse]{
		Message: "Analysis completed",
	})
}
