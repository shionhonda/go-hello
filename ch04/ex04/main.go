package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Before: %v\n", s)
	rotate(s, 5)
	fmt.Printf("After: %v\n", s)
}

// Juggling algorithm
// https://www.geeksforgeeks.org/array-rotation/
func rotate(s []int, d int) {
	d = d % len(s)
	if d == 0 {
		return
	}
	g := gcd(len(s), d)
	for i := 0; i < g; i++ {
		juggle(s, i, d)
	}
}

func juggle(s []int, st int, interval int) {
	var tmp, i, j, k int
	tmp = s[st]
	i = st
	for {
		k = (i + interval) % len(s)
		if k == st {
			break
		}
		j = k
		s[i] = s[j]
		i = j
	}
	s[j] = tmp
}

func gcd(a, b int) int {
	var tmp int
	if a < b {
		a, b = b, a
	}
	for {
		tmp = b
		b = a % b
		a = tmp
		if b == 0 {
			break
		}
	}
	return a
}
