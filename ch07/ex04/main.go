package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

type Reader struct {
	s string
}

func NewReader(s string) *Reader {
	return &Reader{s}
}

func (r *Reader) Read(p []byte) (int, error) {
	if len(r.s) == 0 {
		return 0, io.EOF
	}
	n := copy(p, r.s)
	r.s = r.s[n:]
	return n, nil
}

func main() {
	b, _ := ioutil.ReadFile("fixture/go.html")
	r := NewReader(string(b))
	doc, err := html.Parse(r)
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
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
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
