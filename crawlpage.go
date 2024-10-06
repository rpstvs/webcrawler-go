package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}

	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	stop := cfg.checkPagesLenght()

	if stop {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)

	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse url %s \n", rawCurrentURL)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizeUrl, _ := normalizeURL(rawCurrentURL)

	isFirst := cfg.addPageVisit(normalizeUrl)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s \n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)

	if err != nil {
		fmt.Printf("Error getting the html: %s \n", err)
	}
	urls, _ := getUrlFromHtml(html, cfg.baseURL)

	for _, nexturl := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(nexturl)
	}
}

func (cfg *config) checkPagesLenght() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if len(cfg.pages) > cfg.maxPages {
		return true
	}
	return false
}
