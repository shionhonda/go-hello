package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"go-hello/ch02/ex05/popcount"
)

func bench(f func(uint64) int, x uint64, n int) float64 {
	start := time.Now()
	for i := 0; i < n; i++ {
		f(x)
	}
	secs := time.Since(start).Seconds()
	return secs
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "You must specify at least one argument")
		os.Exit(1)
	}
	arg := os.Args[1]
	x, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot convert to integer: %s", arg)
	}
	n_trials := 1000000
	t := bench(popcount.PopCount, uint64(x), n_trials)
	fmt.Printf("Base: %.6fs\n", t)

	t = bench(popcount.PopCountTrick, uint64(x), n_trials)
	fmt.Printf("Trick: %.6fs\n", t)
}
