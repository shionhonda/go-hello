package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			s := reverseUTF8(sc.Text())
			fmt.Println(s)
		}
	}
}

func reverseBytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseUTF8(s string) string {
	b := []byte(s)
	// reverse bytes inside character
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverseBytes(b[i : i+size])
		i += size
	}
	// reverese bytes across characters
	reverseBytes(b)
	return string(b)
}
