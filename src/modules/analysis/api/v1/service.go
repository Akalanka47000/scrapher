package v1

import (
	"net/url"
	"scrapher/src/modules/analysis/api/v1/dto"
	rodext "scrapher/src/pkg/rod"
	"scrapher/src/utils"
	"sync"

	"github.com/go-rod/rod"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
)

func AnalyseWebPage(targetUrl string) dto.AnalyseWebpageResult {
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
				baseURL := lo.FromPtr(lo.Ok(url.Parse(p.MustInfo().URL)))

				panics := []any{}

				onLink := func(a *rod.Element) {
					href := lo.Ok(a.Property("href")).String()

					external, err := utils.IsExternalLink(href, baseURL)

					if err != nil {
						result.InaccessibleLinkCount++
					} else {
						if external {
							result.ExternalLinkCount++
						} else {
							result.InternalLinkCount++
							href = baseURL.ResolveReference(lo.Ok(url.Parse(href))).String()
						}
					}

					page := pp.MustGet(func() *rod.Page {
						return b.MustPage("")
					})

					err = page.Navigate(href)

					if err != nil {
						log.Warnw("Error visiting link", "link", href, "error", err)
						result.InaccessibleLinkCount++
					}

					pp.Put(page)
				}

				onPanic := func(err any) {
					panics = append(panics, err)
				}

				wg := sync.WaitGroup{}

				for _, a := range lo.Ok(p.Elements("a[href]:not([href^=\"mailto:\"]):not([href^=\"tel:\"])")) {
					wg.Add(1)
					go utils.Protect(func() {
						defer wg.Done()
						onLink(a)
					}, onPanic)
				}

				wg.Wait()

				if len(panics) > 0 {
					panic(panics[0])
				}
			}

			rodext.RunWithNewPagePool(5, analyzeLinks)

			return result
		},
		targetUrl,
	)
}
