package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	page, err := FetchURL(args[0])
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		os.Exit(1)
	}

	links, eErr := ExtractLinks(page)
	if eErr != nil {
		fmt.Println("Error extracting links:", eErr)
		os.Exit(1)
	}

	for _, link := range links {
		fmt.Println(link)
	}
}
