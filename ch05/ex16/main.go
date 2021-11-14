package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strJoin(", ", os.Args[1:]...)
	fmt.Println(s)
}

func strJoin(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}
