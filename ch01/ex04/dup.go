package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	srcs := make(map[string][]string)
	files := os.Args[1:]
	for _, arg := range files {
		err := countLines(arg, counts, srcs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s %v\n", n, line, srcs[line])
		}
	}
}

func countLines(fname string, counts map[string]int, srcs map[string][]string) error {
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		return err
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		srcs[line] = append(srcs[line], fname)
	}
	return nil
}
