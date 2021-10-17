package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("fixture/sample.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)
	c := countWords(sc)
	for k, v := range c {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func countWords(in *bufio.Scanner) map[string]int {
	in.Split(bufio.ScanWords)
	counter := make(map[string]int)
	for in.Scan() {
		w := in.Text()
		w = strings.ToLower(w)
		counter[w]++
	}
	return counter
}
