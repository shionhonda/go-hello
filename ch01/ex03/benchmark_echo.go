package main

import (
	"fmt"
	"strings"
	"time"
)

func createStringArray(n int) []string {
	array := make([]string, 0)
	for i := 0; i < n; i++ {
		array = append(array, "a")
	}
	return array
}

func myJoin(array []string, sep string) string {
	s := ""
	for i, arg := range array {
		if i == 0 {
			s += arg
		} else {
			s += sep + arg
		}
	}
	return s
}

func timer(f func([]string, string) string, array []string) float64 {
	start := time.Now()
	_ = f(array, " ")
	secs := time.Since(start).Seconds()
	return secs
}

func main() {
	array := createStringArray(1000)

	secs := timer(myJoin, array)
	fmt.Printf("My Join: %.6f s\n", secs)

	secs = timer(strings.Join, array)
	fmt.Printf("Native Join: %.6f s\n", secs)
}
