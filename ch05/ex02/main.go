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
	for elem, cnt := range countElements(map[string]int{}, doc) {
		fmt.Printf("%s: %d\n", elem, cnt)
	}
}

func countElements(m map[string]int, n *html.Node) map[string]int {
	// nで見つかったリンクをリストに追加
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	// リストの中身を順番にたどる
	if n.FirstChild != nil {
		m = countElements(m, n.FirstChild)
	}
	if n.NextSibling != nil {
		m = countElements(m, n.NextSibling)
	}
	return m
}
