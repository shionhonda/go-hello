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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// depth-first search
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && isLink(n.Data) {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	// traverse vertically
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	// then horizontally
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

func isLink(tag string) bool {
	switch tag {
	case "a":
		return true
	case "img":
		return true
	case "script":
		return true
	case "style":
		return true
	default:
		return false
	}
}
