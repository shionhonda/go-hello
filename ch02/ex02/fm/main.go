package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"go-hello/ch02/ex02/lenconv"
)

func do(str string) {
	l, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fm: %v\n", err)
		os.Exit(1)
	}
	m := lenconv.Meter(l)
	f := lenconv.Feet(l)
	fmt.Printf("%s = %s, %s = %s\n", m, lenconv.MToF(m), f, lenconv.FToM(f))
}

func main() {
	if len(os.Args) >= 2 {
		for _, arg := range os.Args[1:] {
			do(arg)
		}
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for {
			if sc.Scan() {
				do(sc.Text())
			}
		}
	}

}
