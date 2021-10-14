package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	c1 := sha256.Sum256([]byte((os.Args[1])))
	c2 := sha256.Sum256([]byte((os.Args[2])))
	diff := hammingDist(c1, c2)
	fmt.Printf("Hamming distance of SHA256: %d\n", diff)
}

func hammingDist(a, b [32]byte) int {
	diff := 0
	for i := range a {
		diff += popCount(uint64(a[i] ^ b[i]))
	}
	return diff
}

func popCount(x uint64) int {
	if x == 0 {
		return 0
	}
	y := 0
	for {
		x = x & (x - 1)
		y += 1
		if x == 0 {
			break
		}
	}
	return y
}
