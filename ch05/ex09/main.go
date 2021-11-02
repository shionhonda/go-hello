package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`\$\S+`)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			fmt.Println(expand(sc.Text(), strings.ToUpper))
		}
	}
}

// expand replaces each placeholder with the corresponding value given by f
func expand(s string, f func(string) string) string {
	g := func(s string) string {
		return f(s)[1:]
	}
	return r.ReplaceAllStringFunc(s, g)
}
