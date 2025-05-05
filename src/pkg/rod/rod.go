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
	"github.com/samber/lo"
)

var browser *rod.Browser

// GetHeadlessBrowser creates and returns new headless browser instance using the Rod library
// if it doesn't already exist.
func GetHeadlessBrowser() *rod.Browser {
	if browser != nil {
		return browser
	}
	u := launcher.New().Bin(config.Env.ChromePath).Headless(true).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	global.RegisterShutdownHook("rod-browser-close", func() {
		if browser != nil {
			browser.MustClose()
		}
	})
	return browser
}

// NewHeadlessBrowserSession creates a new headless browser session and executes the provided handler function.
// After the handler function completes, the browser session is closed.
// The handler function receives a pointer to the rod.Browser and a pointer to an ExtendedPage.
// The function returns the result of the handler function.
func NewHeadlessBrowserSession[T any](handler func(*rod.Browser, *ExtendedPage) T, initialURL string) T {
	browser := GetHeadlessBrowser()

	log.Infow("Visiting page", "url", initialURL)

	var e proto.NetworkResponseReceived

	page := browser.MustPage("")

	waitForDocumentDownload := page.WaitEvent(&e)

	err := page.Navigate(initialURL)

	if err != nil {
		log.Errorw("Failed to retrieve webpage", "error", err)
		panic(ErrConnectionError)
	}

	waitForDocumentDownload()

	page.MustWaitLoad()

	var contentType string // Straightforward coalesce doesn't work here due to a bug in `proto.NetworkHeaders`

	if val := e.Response.Headers[strings.ToLower(global.HdrContentType)].Raw(); val != nil {
		contentType = lo.FromBytes[string](lo.Cast[[]byte](val))
	} else if val := e.Response.Headers[global.HdrContentType].Raw(); val != nil {
		contentType = lo.FromBytes[string](lo.Cast[[]byte](val))
	}

	if !strings.Contains(contentType, "text/html") {
		log.Errorw("Invalid content type", "content-type", contentType)
		panic(ErrTargetIsNotValidHTML)
	}

	if e.Response.Status < 200 || e.Response.Status >= 300 {
		log.Errorw("Invalid response status", "status", e.Response.Status)
		panic(global.NewExtendedFiberError(
			fiber.NewError(http.StatusUnprocessableEntity, ErrMsgFailedToAnalyzeWebpage),
			RodErrorDetail{
				TargetStatus: e.Response.Status,
				TargetDetail: http.StatusText(e.Response.Status),
			},
		))
	}

	result := handler(browser, &ExtendedPage{page})

	return result
}

// Creates a new page pool with the specified limit and executes the provided function with it.
// Cleans up the page pool after use.
func RunWithNewPagePool(limit int, fn func(rod.Pool[rod.Page])) {
	pp := rod.NewPagePool(limit)
	defer pp.Cleanup(func(p *rod.Page) {
		if err := p.Close(); err != nil {
			log.Warnw("Error closing page", "error", err)
		}
	})
	fn(pp)
}
