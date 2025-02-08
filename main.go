package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
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
		concurrencyControl: make(chan struct{}, 5), // Limit to 5 concurrent goroutines
		wg:                 &sync.WaitGroup{},
	}

	// Start crawling from the base URL
	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL.String())

	// Wait for all goroutines to finish
	cfg.wg.Wait()

	// Print the pages visited
	for page, count := range cfg.pages {
		fmt.Printf("Visited %s %d times\n", page, count)
	}

}
