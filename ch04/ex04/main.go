package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Before: %v\n", s)
	s = rotate(s, 5)
	fmt.Printf("After: %v\n", s)
}

// TODO: juggling algorithm
func rotate(s []int, l int) []int {
	l = l % len(s)
	return append(s[l:], s[:l]...)
}
