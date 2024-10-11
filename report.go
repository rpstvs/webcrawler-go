package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type Page struct {
	url   string
	count int
}

func createReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
==================
REPORT for %s 
==================
`, baseURL)

	sortedPages := sortMap(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s \n", page.count, page.url)
	}

	CreateReportFile(sortedPages, baseURL)

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

func CreateReportFile(pages []Page, baseURL string) {
	file, err := os.Create("report.txt")

	if err != nil {
		log.Fatal("couldnt create file report")
	}

	defer file.Close()

	opening := fmt.Sprintf(`
==================
Report for %s 
==================
`, baseURL)
	file.Write([]byte(opening))
	for _, page := range pages {
		pageContnt := fmt.Sprintf("Found %d internal links to %s \n", page.count, page.url)
		file.Write([]byte(pageContnt))
	}
}
