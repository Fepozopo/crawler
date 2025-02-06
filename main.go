package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %s\n", os.Args[1])
	}

	htmlBody, err := getHTML(os.Args[1])
	if err != nil {
		fmt.Printf("error getting HTML: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("got HTML body:\n%s\n", htmlBody)
	}

	urls, err := getURLsFromHTML(htmlBody, os.Args[1])
	if err != nil {
		fmt.Printf("error getting URLs: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("got URLs: %v\n", urls)
	}
}
