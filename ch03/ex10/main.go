package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println(comma(os.Args[1]))
}

func comma(num string) string {
	var buf bytes.Buffer
	d := len(num) % 3
	for i, n := range num {
		if i > 0 && i%3 == d {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(n))
	}
	return buf.String()
}
