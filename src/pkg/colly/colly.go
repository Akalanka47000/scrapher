// This package creates a wrapper around the Colly collector to handle errors and logging.
package collyext

import (
	"net/http"
	"scrapher/src/global"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CollyErrorDetail struct {
	TargetStatus     int    `json:"target_status"`
	TargetStatusText string `json:"target_status_text"`
}

type EnhancedCollector struct {
	*colly.Collector
}

func NewCollector() *EnhancedCollector {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Infow("Visiting url", "url", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Errorw("Error while scraping target url",
			"url", r.Request.URL.String(),
			"status_code", r.StatusCode,
			"response", string(r.Body),
		)
		panic(global.NewExtendedFiberError(
			fiber.NewError(http.StatusUnprocessableEntity, "Something went wrong while scraping the given target"),
			CollyErrorDetail{
				TargetStatus:     r.StatusCode,
				TargetStatusText: http.StatusText(r.StatusCode),
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
			fiber.NewError(http.StatusUnprocessableEntity, "We couldn't reach the given target"),
			CollyErrorDetail{
				TargetStatusText: "Connection error",
			},
		))
	}

	return err
}
