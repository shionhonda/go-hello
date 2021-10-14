package main

import "fmt"

const (
	KB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12
	PB = 1e15
	EB = 1e18
	ZB = 1e21
	YB = 1e24
)

func main() {

	fmt.Printf("KB=%f\n", KB)
	fmt.Printf("MB=%f\n", MB)
	fmt.Printf("GB=%f\n", GB)
	fmt.Printf("TB=%f\n", TB)
	fmt.Printf("PB=%f\n", PB)
	fmt.Printf("EB=%f\n", EB)
	fmt.Printf("ZB=%f\n", ZB)
	fmt.Printf("YB=%f\n", YB)
}
