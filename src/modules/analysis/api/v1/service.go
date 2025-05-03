package v1

import (
	"scrapher/src/modules/analysis/api/v1/dto"
	"scrapher/src/pkg/colly"
	"scrapher/src/utils"

	"github.com/gocolly/colly"
	"github.com/samber/lo"
)

func performAnalysis(targetUrl string) dto.PerformAnalysisResult {
	var result dto.PerformAnalysisResult

	c := collyext.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		result.HTMLVersion = utils.ExtractHTMLVersion(string(r.Body))
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		result.PageTitle = e.Text
	})

	c.OnHTML("div", func(e *colly.HTMLElement) {
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

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")

		isExternal, err := utils.IsExternalLink(href, lo.FromPtr(e.Request.URL))

		if err != nil || e.Request.Visit(href) != nil {
			result.InaccessibleLinkCount++
			return
		}

		if !isExternal {
			result.InternalLinkCount++
		} else {
			result.ExternalLinkCount++
		}
	})

	c.Visit(targetUrl)

	return result
}
