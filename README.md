# WebCrawler Golang

This is a small CLI tool done in Golang that allows a user to crawl through a domain.

This tool will output a report with the internal links that found within each link that crawls and will present it in the console and to a TXT File.

## Installation

Clone the project

```bash
go build -o webcrawler
```

## Usage

```bash
./webcrawler [website to crawl] [number_max_threads] [max_links_to_be_crawled]
```

Example of usage

```bash
./webcrawler https://example.com 3 25
```
