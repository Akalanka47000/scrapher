package v1

import (
	"scrapher/src/global"
	"scrapher/src/modules/analysis/api/v1/dto"

	"github.com/gofiber/fiber/v2"
)

func AnalyseWebpage(c *fiber.Ctx) error {
	query := new(dto.AnalyseWebpageRequest)
	c.QueryParser(query)
	result := analyseWebPage(query.URL)
	return c.Status(fiber.StatusOK).JSON(global.Response[dto.AnalyseWebpageResult]{
		Data:    &result,
		Message: "Analysis complete",
	})
}
