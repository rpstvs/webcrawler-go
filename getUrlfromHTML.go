package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getUrlFromHtml(htmlBody string, baseUrl *url.URL) ([]string, error) {

	reader := strings.NewReader(htmlBody)
	tree, err := html.Parse(reader)
	if err != nil {
		log.Println("Couldnt read html body")
	}
	var urlsOnHTML []string
	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, p := range n.Attr {
				if p.Key == "href" {
					href, err := url.Parse(p.Val)

					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", p.Val, err)
						continue
					}
					resolvedUrl := baseUrl.ResolveReference(href)
					urlsOnHTML = append(urlsOnHTML, resolvedUrl.String())
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(tree)

	return urlsOnHTML, nil
}
