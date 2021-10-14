package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(comma(os.Args[1]))
}

func comma(num string) string {
	var minus, dot string
	i := strings.LastIndex(num, "-")
	if i >= 0 {
		minus = "-"
		num = num[1:]
	} else {
		minus = ""
	}
	j := strings.LastIndex(num, ".")
	if j >= 0 {
		dot = num[j:]
		num = num[:j]
	} else {
		dot = ""
	}
	var buf bytes.Buffer
	d := len(num) % 3
	for i, n := range num {
		if i > 0 && i%3 == d {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(n))
	}
	return minus + buf.String() + dot
}
