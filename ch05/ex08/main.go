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
	n := ElementByID(doc, os.Args[2])
	if n != nil {
		fmt.Printf("%s found: %v %v\n", os.Args[2], n.Data, n.Attr)
	} else {
		fmt.Printf("%s not found\n", os.Args[2])
	}
}

// ElementByID finds the first element that has the specified id
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, checkID, checkID)
}

func forEachNode(n *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m := forEachNode(c, id, pre, post)
		if m != nil {
			return m
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}
	return nil
}

func checkID(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}
	}
	return false
}
