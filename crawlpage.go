package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)

	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse url %s \n", rawBaseURL)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)

	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse url %s \n", rawCurrentURL)
		return
	}

	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	normalizeUrl, _ := normalizeURL(rawCurrentURL)

	if _, visited := pages[normalizeUrl]; visited {
		pages[normalizeUrl]++
		return
	}

	pages[normalizeUrl] = 1

	fmt.Printf("crawling %s \n", rawCurrentURL)

	html, _ := getHTML(rawCurrentURL)
	fmt.Println(html)
	//falta err

	urls, _ := getUrlFromHtml(html, rawBaseURL)
	fmt.Println(urls)

	for _, nexturl := range urls {
		crawlPage(rawBaseURL, nexturl, pages)
	}
}
