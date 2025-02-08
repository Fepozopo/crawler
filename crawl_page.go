package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

// addPageVisit increments the counter for a normalized URL.
// It returns true if this is the first visit to this page.
// If the page already exists in our map, we simply increment the count and return false.
func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	// Check if we've seen this URL before:
	if cfg.pages[normalizedURL] > 0 {
		cfg.pages[normalizedURL]++
		return false
	}
	// This is the first time we see this URL.
	cfg.pages[normalizedURL] = 1
	return true
}

// crawlPage performs the crawl for a given raw current URL.
// It normalizes the URL, retrieves its HTML, extracts its links,
// and recursively spawns new goroutines for each discovered link.
func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done() // Mark this goroutine as done when the function exits

	// Parse the current URL.
	currentParsed, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("error parsing current URL: %v\n", err)
		return
	}

	// Resolve the rawCurrentURL relative to the base URL.
	absURL := cfg.baseURL.ResolveReference(currentParsed)

	// Check that the resolved URL is on the same domain.
	if cfg.baseURL.Hostname() != absURL.Hostname() {
		return
	}

	// Normalize the absolute URL.
	currentURL := normalizeURL(absURL.String())

	// Filter out fragment identifiers
	if currentParsed.Fragment != "" {
		return
	}

	// Add the page visit and check if it's the first visit.
	if !cfg.addPageVisit(currentURL) {
		return
	}

	// Get the HTML for the absolute URL.
	htmlBody, err := getHTML(absURL.String())
	if err != nil {
		log.Printf("error getting HTML: %v\n", err)
		return
	}
	fmt.Printf("Visiting URL: %s\n", absURL.String())

	// Get all URLs from the response body HTML.
	urls, err := getURLsFromHTML(htmlBody, absURL.String())
	if err != nil {
		log.Printf("error getting URLs: %v\n", err)
		return
	}

	// Recursively crawl each URL on the page.
	for _, link := range urls {
		fmt.Printf("Discovered link: %s\n", link)

		// Normalize the discovered link
		normalizedLink := normalizeURL(link)

		// Check if the link is already visited
		if !cfg.addPageVisit(normalizedLink) {
			continue
		}

		// Limit the number of concurrent goroutines
		cfg.concurrencyControl <- struct{}{} // Block if the channel is full
		cfg.wg.Add(1)                        // Increment the wait group counter

		go func(link string) {
			defer func() { <-cfg.concurrencyControl }() // Release the slot in the channel
			cfg.crawlPage(link)
		}(normalizedLink)
	}
}
