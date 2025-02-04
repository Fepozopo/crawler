package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlStr string) string {
	// Check if the URL has a scheme; if not, prepend "http://"
	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		urlStr = "http://" + urlStr
	}

	// Parse the URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return urlStr // return the original URL if parsing fails
	}

	// Join the hostname and path together
	normalizedURL := fmt.Sprintf("%s%s", parsedURL.Hostname(), parsedURL.Path)

	// Remove any trailing slash
	normalizedURL = strings.TrimSuffix(normalizedURL, "/")

	return normalizedURL
}
