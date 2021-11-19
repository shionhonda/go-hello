package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	for _, s := range os.Args[1:] {
		if IsPalindrome(newMyStrings(s)) {
			fmt.Println(s, "is a palindrome")
		} else {
			fmt.Println(s, "is not a palindrome")
		}
	}
}

type myStrings []string

func newMyStrings(s string) myStrings {
	return strings.Split(s, "")
}

func (s myStrings) Len() int {
	return len(s)
}

func (s myStrings) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s myStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
