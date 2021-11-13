package main

import (
	"fmt"
	"os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	breadthFirst(crawl, os.Args[1])
}

func breadthFirst(f func(item string) []string, origin string) {
	seen := make(map[string]bool)
	worklist := []string{origin}

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				fmt.Println(item)
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(course string) []string {
	list, ok := prereqs[course]
	if !ok {
		return nil
	}
	return list
}
