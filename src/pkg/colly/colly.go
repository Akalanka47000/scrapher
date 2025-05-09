// This package creates a wrapper around the Colly collector to handle errors and logging.
// Currently, it is not used in the codebase, but left here for future use if needed.
package collyext

import (
	"net/http"
	"scrapher/src/global"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type EnhancedCollector struct {
	*colly.Collector
}

func NewCollector() *EnhancedCollector {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Infow("Visiting webpage", "url", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Errorw("Error while scraping webpage",
			"url", r.Request.URL.String(),
			"status_code", r.StatusCode,
			"response", string(r.Body),
		)
		panic(global.NewExtendedFiberError(
			fiber.NewError(http.StatusUnprocessableEntity, ErrFailedToAnalyzeWebpage),
			CollyErrorDetail{
				TargetStatus: r.StatusCode,
				TargetDetail: http.StatusText(r.StatusCode),
			},
		))
	})

	return &EnhancedCollector{c}
}

func (c *EnhancedCollector) Visit(url string) error {
	err := c.Collector.Visit(url)

	if err != nil {
		log.Error("Error connecting to target url: ", err)
		panic(global.NewExtendedFiberError(
			fiber.NewError(http.StatusUnprocessableEntity, ErrFailedToAnalyzeWebpage),
			CollyErrorDetail{
				TargetDetail: "Connection error",
			},
		))
	}

	return err
}
