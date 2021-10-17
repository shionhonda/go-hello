package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const space string = " "

func main() {
	str := "1 2 3  4 \n　5 日本語 7"
	s := []byte(str)
	fmt.Printf("Before: %s\n", string(s))
	s = removeDupSpace(s)
	fmt.Printf("After : %s\n", string(s))
}

func removeDupSpace(s []byte) []byte {
	var i, dups int
	for i = 0; i < len(s); {
		r0, sz0 := utf8.DecodeRune(s[i:])
		r1, sz1 := utf8.DecodeRune(s[i+sz0:])
		if unicode.IsSpace(r0) && unicode.IsSpace(r1) {
			copy(s[i+sz0:], s[i+sz0+sz1:])
			dups += sz1
			// to deal with trailing spaces
			if i+dups >= len(s)-1 {
				return s[:i+1]
			}
		} else {
			// spaces are deduplicated until i+sz0
			i += sz0
		}
	}
	return s[:i-dups]
}
