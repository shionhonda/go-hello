package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, start, end)
	return nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		startElement(n)
	case html.TextNode:
		startText(n)
	case html.CommentNode:
		startComment(n)
	}
}

func end(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		endElement(n)
	}
}

func startElement(n *html.Node) {
	attrs := []string{}
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf("%s=%s", a.Key, a.Val))
	}
	var attrStr string
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	if n.FirstChild == nil {
		fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrStr)
	} else {
		fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrStr)
	}
	depth++
}

func endElement(n *html.Node) {
	depth--
	if n.FirstChild == nil {
		return
	}
	fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
}

func startText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	fmt.Printf("%*s%s\n", depth*2, "", text)
}

func startComment(n *html.Node) {
	fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
}
