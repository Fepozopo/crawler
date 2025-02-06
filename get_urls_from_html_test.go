package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	htmlBody := `
		<html>
			<body>
				<a href="http://example.com">Example</a>
				<a href="/about">About</a>
				<a href="https://example.com/contact">Contact</a>
				<a href="relative/path">Relative</a>
			</body>
		</html>
	`
	rawBaseURL := "http://example.com"

	expectedURLs := []string{
		"http://example.com",
		"http://example.com/about",
		"https://example.com/contact",
		"http://example.com/relative/path",
	}

	urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(urls, expectedURLs) {
		t.Errorf("expected %v, got %v", expectedURLs, urls)
	}
}

func TestGetURLsFromHTML_EmptyBaseURL(t *testing.T) {
	htmlBody := `<html><body><a href="/about">About</a></body></html>`
	_, err := getURLsFromHTML(htmlBody, "")
	if err == nil {
		t.Fatal("expected an error for empty base URL, got none")
	}
}
