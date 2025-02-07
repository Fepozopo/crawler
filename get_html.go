package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// Make the HTTP GET request
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("bad HTTP status code: %v", resp.StatusCode)
	}

	// Check if the content-type header starts with text/html
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("content type is not text/html: %v", contentType)
	}

	// Read the response body
	htmlBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(htmlBody), nil
}
