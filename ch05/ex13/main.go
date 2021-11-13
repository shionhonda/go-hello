package main

import (
	"fmt"
	"go-hello/ch05/ex13/links"
	"log"
	"net/url"
	"os"
)

func main() {
	breadthFirst(crawl, os.Args[1])
}

func breadthFirst(f func(item string) []string, origin string) {
	seen := make(map[string]bool)
	hostname, err := getHostname(origin)
	if err != nil {
		log.Fatal(err)
	}
	worklist := []string{origin}

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				h, err := getHostname(item)
				if err == nil {
					if h == hostname {
						fmt.Println(item)
					} else {
						fmt.Printf("skipping: %s\n", item)
					}
				}
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(urlStr string) []string {
	list, err := links.Extract(urlStr)
	if err != nil {
		log.Print(err)
	}
	return list
}

func getHostname(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}
