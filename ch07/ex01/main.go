package main

import (
	"bufio"
	"fmt"
	"go-hello/ch07/ex01/counter"
	"os"
)

func main() {
	var wc counter.WordCounter
	var lc counter.LineCounter

	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			text := sc.Text()
			wc.Write([]byte(text))
			lc.Write([]byte(text))
		} else {
			break
		}
	}

	fmt.Printf("Words: %d\n", wc)
	fmt.Printf("Lines: %d\n", lc)
}
