package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// Convert os.Args[2] and os.Args[3] to integers if they exist
	var concurrency, maxPages int
	if len(os.Args) > 2 {
		var err error
		concurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("error parsing concurrency: %v\n", err)
			os.Exit(1)
		}
	}
	if len(os.Args) > 3 {
		var err error
		maxPages, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("error parsing max pages: %v\n", err)
			os.Exit(1)
		}
	}

	// Set concurrency to 4 if it's not provided or less than 1
	if concurrency <= 0 {
		concurrency = 4
	}

	// Set maxPages to 0 if it's not provided or less than 1
	if maxPages <= 0 {
		maxPages = 0
	}

	// Prepend "http://" to the website if it's missing
	if !strings.HasPrefix(os.Args[1], "http://") && !strings.HasPrefix(os.Args[1], "https://") {
		os.Args[1] = "http://" + os.Args[1]
	}

	// Parse the base URL
	baseURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("error parsing base URL: %v\n", err)
		os.Exit(1)
	}

	// Print the starting message
	fmt.Printf("Starting crawl of: %s\n", baseURL.String())

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, concurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	// Start crawling from the base URL
	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL.String())

	// Wait for all goroutines to finish
	cfg.wg.Wait()

	// Print the report
	printReport(cfg.pages, baseURL.String())

}
