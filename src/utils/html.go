package utils

import (
	"net/url"
	"strings"
)

func ExtractHTMLVersion(html string) string {
	lowercaseHTML := strings.ToLower(html)
	switch {
	case strings.HasPrefix(lowercaseHTML, "<!doctype html>"):
		return "HTML5"
	case strings.Contains(lowercaseHTML, "xhtml"):
		return "XHTML"
	case strings.Contains(lowercaseHTML, "html 4.01 transitional"):
		return "HTML 4.01 Transitional"
	case strings.Contains(lowercaseHTML, "html 4.01 strict"):
		return "HTML 4.01 Strict"
	case strings.Contains(lowercaseHTML, "html 4.01 frameset"):
		return "HTML 4.01 Frameset"
	case strings.Contains(lowercaseHTML, "html 4.0 transitional"):
		return "HTML 4.0 Transitional"
	case strings.Contains(lowercaseHTML, "html 4.0 strict"):
		return "HTML 4.0 Strict"
	case strings.Contains(lowercaseHTML, "html 4.0 frameset"):
		return "HTML 4.0 Frameset"
	case strings.Contains(lowercaseHTML, "html 3.2"):
		return "HTML 3.2"
	default:
		return "Unknown"
	}
}

func IsExternalLink(link string, source url.URL) (bool, error) {
	parsed, err := url.Parse(link)
	if err != nil {
		return false, err
	}

	if parsed.Host == source.Host || parsed.Host == "" {
		return false, nil
	}
	return true, nil
}
