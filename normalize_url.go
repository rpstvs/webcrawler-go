package main

import "net/url"

func normalizeURL(urL string) (string, error) {

	parsedURL, err := url.Parse(urL)

	if err != nil {
		return "", err
	}
	return parsedURL.Host + parsedURL.Path, nil
}
