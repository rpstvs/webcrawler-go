package main

import (
	"fmt"
	"sort"
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
	pagesSlice := []Page{}

	for url, count := range pages {
		pagesSlice = append(pagesSlice, Page{url: url, count: count})
	}

	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].count == pagesSlice[j].count {
			return pagesSlice[i].url < pagesSlice[j].url
		}
		return pagesSlice[i].count > pagesSlice[j].count
	})

	return pagesSlice

}
