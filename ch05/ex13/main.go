package main

import (
	"fmt"
	"go-hello/ch05/ex13/links"
	"log"
)

func main() {
	list, _ := links.Extract("https://www.gopl.io/")
	for _, link := range list {
		fmt.Println(link)
	}
	// breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
