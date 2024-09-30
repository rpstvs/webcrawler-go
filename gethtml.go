package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)

	if err != nil {
		return "", fmt.Errorf("error on http")
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		fmt.Println("Error on response")
		return "", fmt.Errorf("Error on response")
	}

	contentType := resp.Header.Get("Content-Type")

	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("error content type")
	}

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("Error on reading body")
	}

	c := string(dat)

	fmt.Println(c)

	return c, nil
}
