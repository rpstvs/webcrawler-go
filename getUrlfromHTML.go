package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func getUrlFromHtml(htmlBody, rawBaseURL string) ([]string, error) {
	reader := strings.NewReader(htmlBody)

	tree, err := html.Parse(reader)
	if err != nil {
		log.Println("Couldnt read html body")
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		fmt.Println(n.Data)
		if n.Type == html.ElementNode && n.Data == "a" {

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(tree)

	return []string{""}, nil
}
