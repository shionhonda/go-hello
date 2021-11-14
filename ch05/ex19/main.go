package main

import "fmt"

func main() {
	msg, err := dontPanic()
	fmt.Println(msg)
	fmt.Println(err)
}

func dontPanic() (msg string, err error) {
	defer func() {
		if p := recover(); p != nil {
			msg = "hello"
			err = fmt.Errorf("recovered: %s", p.(string))
		}
	}()
	panic("try panic")
}
