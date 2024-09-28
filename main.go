package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no website provided")
	} else if len(os.Args) > 2 {
		log.Fatal("too many arguments")
	} else {
		baseURL := os.Args[1]
		fmt.Printf("Starting webcrawler on: %s \n", baseURL)
	}

}
