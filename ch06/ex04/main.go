package main

import (
	"fmt"
	"go-hello/ch06/ex02/intset"
)

func main() {
	var x, y intset.IntSet
	x.AddAll(1, 9, 144)
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
