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

func analyseWebPage(targetUrl string) dto.AnalyseWebpageResult {
	return rodext.NewHeadlessBrowserSession(
		func(b *rod.Browser, p *rodext.ExtendedPage) (result dto.AnalyseWebpageResult) {
			result.HTMLVersion = p.HTMLVersion()
			result.PageTitle = p.MustInfo().Title
			result.HeadingCounts.H1 = p.ElementCount("h1")
			result.HeadingCounts.H2 = p.ElementCount("h2")
			result.HeadingCounts.H3 = p.ElementCount("h3")
			result.HeadingCounts.H4 = p.ElementCount("h4")
			result.HeadingCounts.H5 = p.ElementCount("h5")
			result.HeadingCounts.H6 = p.ElementCount("h6")

			result.ContainsLoginForm = p.ContainsLoginForm()

			analyzeLinks := func(pp rod.Pool[rod.Page]) {
				wg := sync.WaitGroup{}

				baseURL := lo.FromPtr(lo.Ok(url.Parse(p.MustInfo().URL)))

				for _, a := range lo.Ok(p.Elements("a[href]")) {
					wg.Add(1)
					go func() {
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

						page, err := pp.Get(func() (*rod.Page, error) {
							return b.Page(proto.TargetCreateTarget{URL: href})
						})

						if err == nil {
							err = page.WaitLoad()
						}

						if err != nil {
							log.Warnw("Error visiting link", "link", href, "error", err)
							result.InaccessibleLinkCount++
						}

						pp.Put(page)
					}()
				}

				wg.Wait()
			}

			rodext.RunWithNewPagePool(3, analyzeLinks)

			return result
		},
		targetUrl,
	)
}
