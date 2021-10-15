package main

import (
	"fmt"
	"unicode"
)

const space string = " "

func main() {
	str := "1 2 3  4 　5 6 7"
	s := []byte(str)
	fmt.Printf("Before: %v\n", s)
	s = removeDupSpace(s)
	fmt.Printf("After : %v\n", s)
}

func removeDupSpace(s []byte) []byte {
	dups := 0
	for i := 0; i < len(s)-1; i++ {
		if i+dups >= len(s)-1 {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			if i < len(s)-2 {
				copy(s[i+1:], s[i+2:])
				// s[i+1] = byte(space[0])
			}
			dups++
		}
	}
	return s[:len(s)-dups]
}
