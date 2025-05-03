package v1

import (
	"scrapher/src/modules/analysis/api/v1/dto"
	"scrapher/src/pkg/colly"

	"github.com/gocolly/colly"
)

func performAnalysis(targetUrl string) dto.PerformAnalysisResult {
	var result dto.PerformAnalysisResult

	c := collyext.NewCollector()

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		result.HeadingCounts.H1++
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		result.HeadingCounts.H2++
	})

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		result.HeadingCounts.H3++
	})

	c.OnHTML("h4", func(e *colly.HTMLElement) {
		result.HeadingCounts.H4++
	})

	c.OnHTML("h5", func(e *colly.HTMLElement) {
		result.HeadingCounts.H5++
	})

	c.OnHTML("h6", func(e *colly.HTMLElement) {
		result.HeadingCounts.H6++
	})

	c.Visit(targetUrl)

	return result
}
