package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\n")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Printf("=============================\n")

	// Create a slice of structs to hold the pages and their counts
	type pageCount struct {
		count int
		page  string
	}

	pageCounts := make([]pageCount, 0, len(pages))
	for page, count := range pages {
		pageCounts = append(pageCounts, pageCount{count, page})
	}

	// Sort the slice of page counts
	sort.Slice(pageCounts, func(i, j int) bool {
		if pageCounts[i].count == pageCounts[j].count {
			return pageCounts[i].page < pageCounts[j].page
		}
		return pageCounts[i].count > pageCounts[j].count
	})

	// Print the page counts
	for _, pc := range pageCounts {
		fmt.Printf("Found %d internal links to %s\n", pc.count, pc.page)
	}
}
