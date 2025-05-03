// This package creates a wrapper around Rod to make it easier to use.
package rodext

import (
	"scrapher/src/utils"

	"github.com/go-rod/rod"
)

type ExtendedPage struct {
	*rod.Page
}

// HTML of the page including the doctype.
func (p ExtendedPage) RawHTML() string {
	return p.MustEval(`() => {
			return new XMLSerializer().serializeToString(document.doctype);
		}
	`).String()
}

// Extracts the HTML version from the page's raw HTML.
func (p ExtendedPage) HTMLVersion() string {
	return utils.ExtractHTMLVersion(p.RawHTML())
}
