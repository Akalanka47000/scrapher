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
			try {
				return new XMLSerializer().serializeToString(document.doctype);
			} catch (e) {
				return "" 
			}
		}
	`).String()
}

// Extracts the HTML version from the page's raw HTML.
func (p ExtendedPage) HTMLVersion() string {
	return utils.ExtractHTMLVersion(p.RawHTML())
}

// Retrieves the number of elements with the specified selector.
// Efficient than getting the length of `p.Elements(selector)` as it doesn't
// load the elements themselves into memory.
func (p ExtendedPage) ElementCount(selector string) int {
	return p.MustEval(`(selector) => {
		return document.querySelectorAll(selector).length;
	}`, selector).Int()
}

// Checks if the page contains a login form by looking at a few common patterns.
// This most likely is not 100% accurate and it can't pick up forms which are injected
// dynamically by a user action (e.g. clicking a button).
func (p ExtendedPage) ContainsLoginForm() bool {
	return p.MustEval(`() => [...document.forms].some(form => {
		return [...form.elements].some(e => e.type === 'password') ||
			/login|sign.?in|auth/i.test(form.outerHTML)
	})`).Bool()
}
