package main

import (
	"fmt"
	"io"
	"net/http"
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

	// Check the response content type
	if resp.Header.Get("Content-Type") != "text/html" {
		return "", fmt.Errorf("bad content type: %v", resp.Header.Get("Content-Type"))
	}

	// Read the response body
	htmlBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(htmlBody), nil
}
