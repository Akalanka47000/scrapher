package v1

import (
	"scrapher/src/global"
	"scrapher/src/modules/analysis/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func PerformAnalysis(c *fiber.Ctx) error {
	payload := new(dto.PerformAnalysisRequest)
	c.BodyParser(payload)
	result := performAnalysis(payload.TargetURL)
	return c.Status(fiber.StatusOK).JSON(global.Response[dto.PerformAnalysisResult]{
		Data:    &result,
		Message: "Analysis completed",
	})
}
