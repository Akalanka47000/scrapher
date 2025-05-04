// This package creates a wrapper around Rod to make it easier to use.
package rodext

import (
	"net/http"
	"scrapher/src/config"
	"scrapher/src/global"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// NewHeadlessBrowser creates a new headless browser instance using the Rod library.
func NewHeadlessBrowser() *rod.Browser {
	u := launcher.New().Bin(config.Env.ChromePath).Headless(true).MustLaunch()
	return rod.New().ControlURL(u)
}

// NewHeadlessBrowserSession creates a new headless browser session and executes the provided handler function.
// After the handler function completes, the browser session is closed.
// The handler function receives a pointer to the rod.Browser and a pointer to an ExtendedPage.
// The function returns the result of the handler function.
func NewHeadlessBrowserSession[T any](handler func(*rod.Browser, *ExtendedPage) T, initialURL string) T {
	browser := NewHeadlessBrowser()

	browser.MustConnect()
	defer browser.MustClose()

	log.Infow("Visiting target", "url", initialURL)

	var e proto.NetworkResponseReceived

	page := browser.MustPage("")

	waitForDocumentDownload := page.WaitEvent(&e)

	err := page.Navigate(initialURL)

	if err != nil {
		log.Errorw("Failed to retrieve webpage", "error", err)
		panic(global.NewExtendedFiberError(
			fiber.NewError(http.StatusUnprocessableEntity, ErrFailedToAnalyzeTargetURL),
			RodErrorDetail{
				TargetDetail: ErrDetailConnectionError,
			},
		))
	}

	waitForDocumentDownload()

	page.MustWaitStable()

	contentType := e.Response.Headers[strings.ToLower(global.HdrContentType)].Str()

	if !strings.Contains(contentType, "text/html") {
		log.Errorw("Invalid content type", "content-type", contentType)
		panic(global.NewExtendedFiberError(
			fiber.NewError(http.StatusUnprocessableEntity, ErrFailedToAnalyzeTargetURL),
			RodErrorDetail{
				TargetDetail: ErrDetailTargetUrlIsNotValidHTML,
			},
		))
	}

	if e.Response.Status < 200 || e.Response.Status >= 300 {
		log.Errorw("Invalid response status", "status", e.Response.Status)
		panic(global.NewExtendedFiberError(
			fiber.NewError(http.StatusUnprocessableEntity, ErrFailedToAnalyzeTargetURL),
			RodErrorDetail{
				TargetStatus: e.Response.Status,
				TargetDetail: http.StatusText(e.Response.Status),
			},
		))
	}

	result := handler(browser, &ExtendedPage{page})

	return result
}
