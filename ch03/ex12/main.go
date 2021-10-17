package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b := isAnagram(os.Args[1], os.Args[2])
	if b {
		fmt.Printf("%s and %s are anagram of each other.\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s and %s are not anagram.\n", os.Args[1], os.Args[2])

	}
}

func isAnagram(s0, s1 string) bool {
	// s1 にしかない文字があったらfalse
	if len(s0) != len(s1) {
		return false
	}
	var pool, s string
	for _, r := range s0 {
		s = string(r)
		if !strings.Contains(pool, s) {
			if strings.Count(s0, s) != strings.Count(s1, s) {
				return false
			}
			pool += s
		}
	}
	return true
}
