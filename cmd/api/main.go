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
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(page)
}
