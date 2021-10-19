package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for i, s := range findTexts(nil, doc) {
		fmt.Printf("%03d: %s\n", i, s)
	}
}

func findTexts(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	// traverse vertically
	if n.FirstChild != nil && isShown(n.FirstChild) {
		texts = findTexts(texts, n.FirstChild)
	}
	// then horizontally
	if n.NextSibling != nil && isShown(n.NextSibling) {
		texts = findTexts(texts, n.NextSibling)
	}
	return texts
}

func isShown(n *html.Node) bool {
	return n.Data != "script" && n.Data != "style"
}
