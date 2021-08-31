package main

import (
	"os"
	"time"
)

func myJoin(list []string) string {
	s, sep := "", ""
	for _, arg := range list {
		s += sep + arg
		sep = " "
	}
	return s
}

func main() {
	start := time.Now()
	s := myJoin(os.Args)
	secs := time.Since(start).Seconds()
}
