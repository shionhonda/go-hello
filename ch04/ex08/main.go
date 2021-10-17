package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	c := countChar(in)
	for k, v := range c {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func countChar(in *bufio.Reader) map[string]int {
	counter := make(map[string]int)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(r) {
			counter["letter"] += 1
		} else if unicode.IsNumber(r) {
			counter["number"] += 1
		} else if unicode.IsSpace(r) {
			counter["space"] += 1
		}
	}
	return counter
}
