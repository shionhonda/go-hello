package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ns := ElementsByTagName(doc, os.Args[2:]...)
	for _, n := range ns {
		fmt.Println(n.Data)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var result []*html.Node
	forEachNode(doc, &result, checkName, name)
	return result
}

func forEachNode(n *html.Node, ns *[]*html.Node, check func(*html.Node, []string) bool, names []string) {
	if check != nil {
		if check(n, names) {
			*ns = append(*ns, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, ns, check, names)
	}
}

func checkName(n *html.Node, names []string) bool {
	for _, name := range names {
		if n.Data == name {
			return true
		}
	}
	return false
}
