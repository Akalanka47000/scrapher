package utils

import (
	"net/url"
)

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
