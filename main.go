package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {

		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		baseURL := os.Args[1]
		c, _ := getHtmml(baseURL)
		fmt.Printf("starting crawl: %s \n", baseURL)
		fmt.Println(c)
	}

}
