package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urL string) (string, error) {

	parsedURL, err := url.Parse(urL)

	if err != nil {
		return "", err
	}

	normalizedURL := parsedURL.Host + parsedURL.Path

	normalizedURL = strings.ToLower(normalizedURL)

	normalizedURL = strings.TrimSuffix(normalizedURL, "/")

	return normalizedURL, nil
}
