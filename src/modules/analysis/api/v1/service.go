package v1

import (
	"github.com/go-rod/rod"
	"scrapher/src/modules/analysis/api/v1/dto"
	rodext "scrapher/src/pkg/rod"
)

func performAnalysis(targetUrl string) dto.PerformAnalysisResult {
	return rodext.NewHeadlessBrowserSession(
		func(b *rod.Browser, p *rodext.ExtendedPage) (result dto.PerformAnalysisResult) {
			result.HTMLVersion = p.HTMLVersion()

			return result
		},
		targetUrl,
	)
}
