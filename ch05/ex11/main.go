package main

import (
	"fmt"
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
	"linear algebra":        {"calculus"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Println(err)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	var stack []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				stack = append(stack, item)
				for _, prereq := range m[item] {
					if contains(stack, prereq) {
						return fmt.Errorf("cycle detected")
					}
				}
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				order = append(order, item)
				stack = stack[:len(stack)-1]
			}
		}
		return nil
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	err := visitAll(keys)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
