package main

import (
	"bufio"
	"fmt"
	"go-hello/ch07/ex14/eval"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			expr, err := eval.Parse(sc.Text())
			if err != nil {
				fmt.Printf("parse error: %v\n", err)
				continue
			}
			fmt.Println(expr.Eval(eval.Env{}))
		}
	}

}
