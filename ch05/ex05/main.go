package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatalf("countWordsAndImages: %v\n", err)
	}
	fmt.Printf("Words: %d, Images: %d\n", words, images)
}

// CountWordsAndImages throws GET request to url
// and returns the count of words and images found in the HTML document
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	return countRecursive(0, 0, doc)
}

// depth-first search
func countRecursive(words, images int, n *html.Node) (int, int) {
	if n.Type == html.TextNode {
		ws := strings.Fields(n.Data)
		words += len(ws)
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	// traverse vertically
	if n.FirstChild != nil {
		words, images = countRecursive(words, images, n.FirstChild)
	}
	// then horizontally
	if n.NextSibling != nil {
		words, images = countRecursive(words, images, n.NextSibling)
	}
	return words, images
}
