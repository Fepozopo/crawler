package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// Print the starting message
	fmt.Printf("starting crawl of: %s\n", os.Args[1])

	// Prepend "http://" to the website if it's missing
	if !strings.HasPrefix(os.Args[1], "http://") && !strings.HasPrefix(os.Args[1], "https://") {
		os.Args[1] = "http://" + os.Args[1]
	}

	// Crawl the website and when it's complete,
	// print the keys and values of the pages map
	pages := make(map[string]int)
	crawlPage(os.Args[1], os.Args[1], pages)
	for k, v := range pages {
		fmt.Printf("%v: %v\n", k, v)
	}
}
