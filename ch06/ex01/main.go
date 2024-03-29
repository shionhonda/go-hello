package main

import (
	"fmt"
	"go-hello/ch06/ex01/intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y = *x.Copy()
	fmt.Println(y.String())

	x.Remove(1)
	fmt.Println(x.Len())
	fmt.Println(y.Len())

	x.Clear()
	fmt.Println(x.Len())
	fmt.Println(y.Len())
}
