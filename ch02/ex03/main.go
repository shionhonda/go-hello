package main

import (
	"fmt"
	"os"
	"strconv"

	"go-hello/ch02/ex03/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot convert to integer: %s", arg)
		}
		c := popcount.PopCountLoop(uint64(x))
		fmt.Printf("%d\n", c)
	}
}
