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

		cfg, err := configure(baseURL, 3)

		if err != nil {
			fmt.Printf("error configure %v: ", err)
			return
		}

		fmt.Printf("starting crawl: %s \n", baseURL)

		cfg.wg.Add(1)
		go cfg.crawlPage(baseURL)
		cfg.wg.Wait()
		fmt.Println(cfg.pages)
	}

}
