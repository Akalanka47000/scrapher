package utils

import (
	"net/url"
	"strings"
)

func ExtractHTMLVersion(html string) string {
	lowercaseHTML := strings.ToLower(html)
	if strings.HasPrefix(lowercaseHTML, "<!doctype html>") {
		return "HTML5"
	} else if strings.Contains(lowercaseHTML, "xhtml") {
		return "XHTML"
	} else if strings.Contains(lowercaseHTML, "transitional") {
		return "HTML 4.01 Transitional"
	}
	return "Unknown"
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
