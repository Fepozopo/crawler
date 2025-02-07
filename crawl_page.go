package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Parse the base URL once as a reference
	baseParsed, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Printf("error parsing base URL: %v\n", err)
		return
	}

	// Parse the current URL.
	currentParsed, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("error parsing current URL: %v\n", err)
		return
	}

	// Resolve the rawCurrentURL relative to the base URL.
	absURL := baseParsed.ResolveReference(currentParsed)

	// Check that the resolved URL is on the same domain.
	if baseParsed.Hostname() != absURL.Hostname() {
		return
	}

	// Normalize the absolute URL.
	currentURL := normalizeURL(absURL.String())

	// If the pages map already has an entry for normalized current URL, increment and return.
	if pages[currentURL] > 0 {
		pages[currentURL]++
		return
	}

	// Otherwise, add an entry to the pages map for the normalized current URL, and set count = 1.
	pages[currentURL] = 1

	// Get the HTML for the absolute URL.
	htmlBody, err := getHTML(absURL.String())
	if err != nil {
		log.Printf("error getting HTML: %v\n", err)
		return
	}
	fmt.Printf("crawling: %s\n", currentURL)

	// Get all URLs from the response body HTML.
	urls, err := getURLsFromHTML(htmlBody, absURL.String())
	if err != nil {
		log.Printf("error getting URLs: %v\n", err)
		return
	}

	// Recursively crawl each URL on the page.
	for _, link := range urls {
		fmt.Printf("crawling: %s\n", link)
		crawlPage(rawBaseURL, link, pages)
	}

}
