package main

import (
	"fmt"
)

type Page struct {
	url   string
	count int
}

func createReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
	==================
	Report for %s 
	==================`, baseURL)

	sortedPages := sortMap(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s \n", page.count, page.url)
	}

}

func sortMap(pages map[string]int) []Page {

	return []Page{}
}
