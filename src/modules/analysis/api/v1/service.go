package v1

import (
	"net/url"
	"scrapher/src/modules/analysis/api/v1/dto"
	rodext "scrapher/src/pkg/rod"
	"scrapher/src/utils"
	"sync"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
)

func performAnalysis(targetUrl string) dto.PerformAnalysisResult {
	return rodext.NewHeadlessBrowserSession(
		func(b *rod.Browser, p *rodext.ExtendedPage) (result dto.PerformAnalysisResult) {
			result.HTMLVersion = p.HTMLVersion()
			result.PageTitle = p.MustInfo().Title
			result.HeadingCounts.H1 = p.ElementCount("h1")
			result.HeadingCounts.H2 = p.ElementCount("h2")
			result.HeadingCounts.H3 = p.ElementCount("h3")
			result.HeadingCounts.H4 = p.ElementCount("h4")
			result.HeadingCounts.H5 = p.ElementCount("h5")
			result.HeadingCounts.H6 = p.ElementCount("h6")

			result.ContainsLoginForm = p.ContainsLoginForm()

			wg := sync.WaitGroup{}

			baseURL := lo.FromPtr(lo.Ok(url.Parse(p.MustInfo().URL)))

			for _, a := range lo.Ok(p.Elements("a[href]")) {
				wg.Add(1)
				go func(a *rod.Element) {
					defer wg.Done()

					href := lo.Ok(a.Property("href")).String()

					external, err := utils.IsExternalLink(href, baseURL)

					if err != nil {
						result.InaccessibleLinkCount++
					}

					if external {
						result.ExternalLinkCount++
					} else {
						result.InternalLinkCount++
						href = baseURL.ResolveReference(lo.Ok(url.Parse(href))).String()
					}

					tab, err := b.Page(proto.TargetCreateTarget{URL: href, Background: true})
					if err == nil {
						err = tab.WaitLoad()
					}

					if err != nil {
						log.Warnw("Error visiting link", "link", href, "error", err)
						result.InaccessibleLinkCount++
					}

					tab.Close()
				}(a)
			}

			wg.Wait()

			return result
		},
		targetUrl,
	)
}
