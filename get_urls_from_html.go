package main

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// getURLsFromHTML extracts all URLs from <a> tags in the provided HTML body
// and converts relative URLs to absolute URLs based on the rawBaseURL.
func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	if rawBaseURL == "" {
		return nil, errors.New("base URL cannot be empty")
	}

	// Parse the base URL
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	var urls []string

	// Parse the HTML
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	// Traverse the HTML nodes
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					var absoluteURL string
					// Check if the URL is absolute
					if strings.HasPrefix(attr.Val, "http://") || strings.HasPrefix(attr.Val, "https://") {
						absoluteURL = attr.Val
					} else {
						// Resolve the relative URL
						absoluteURL = base.ResolveReference(&url.URL{Path: attr.Val}).String()
					}
					urls = append(urls, absoluteURL)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return urls, nil
}
