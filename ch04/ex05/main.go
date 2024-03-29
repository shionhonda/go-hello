package main

import "fmt"

func main() {
	s := []string{"1", "2", "2", "3", "4", "5"}
	fmt.Printf("Before: %v\n", s)
	s = removeNextDups(s)
	fmt.Printf("After: %v\n", s)
}

func removeNextDups(s []string) []string {
	var j, dups int
	for i := 0; i < len(s)-1; i++ {
		if i+dups >= len(s)-1 {
			break
		}
		if s[i] == s[i+1] {
			if i < len(s)-2 {
				copy(s[i+1:], s[i+2:])
			}
			dups++
		} else {
			// elements are deduplicated until j
			j = i + 1
		}
	}
	return s[:j+1]
}
