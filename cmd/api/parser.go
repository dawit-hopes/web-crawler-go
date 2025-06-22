package main

import (
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinks(htmlString string) ([]string, error) {
	text := strings.NewReader(htmlString)

	doc, err := html.Parse(text)

	if err != nil {
		return []string{}, err
	}

	var links []string
	visit(doc, &links)
	return links, nil
}

func visit(node *html.Node, links *[]string) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attar := range node.Attr {
			if attar.Key == "href" && strings.HasPrefix(attar.Val, "http") {
				*links = append(*links, attar.Val)
			}
		}
	}

	if node.FirstChild != nil {
		visit(node.FirstChild, links)
	}

	if node.NextSibling != nil {
		visit(node.NextSibling, links)
	}
}
