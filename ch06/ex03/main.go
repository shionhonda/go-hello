package main

import (
	"fmt"
	"go-hello/ch06/ex03/intset"
)

func main() {
	var x, y intset.IntSet
	x.AddAll(1, 2, 3, 4, 5)
	y.AddAll(1, 3, 5, 7, 9)
	x.IntersectWith(&y)
	fmt.Println(x.String())

	x.AddAll(1, 2, 3, 4, 5)
	x.DifferenceWith(&y)
	fmt.Println(x.String())

	x.AddAll(1, 2, 3, 4, 5)
	x.SymmetricDifference(&y)
	fmt.Println(x.String())
}
