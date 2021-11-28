package main

import (
	"bufio"
	"fmt"
	"go-hello/ch07/ex15/eval"
	"os"
	"strconv"
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
			vars, err := parseVars(expr)
			env := ask(sc, vars)
			fmt.Println(expr.Eval(env))
		}
	}

}

func ask(sc *bufio.Scanner, vars map[eval.Var]bool) eval.Env {
	var env = make(eval.Env)
	for k := range vars {
		if _, ok := env[k]; !ok {
			for {
				fmt.Printf("%s = ?", k)
				if sc.Scan() {
					s := sc.Text()
					f, err := strconv.ParseFloat(s, 64)
					if err != nil {
						continue
					}
					env[k] = f
					break
				}
			}

		}
	}
	return env
}

func parseVars(expr eval.Expr) (map[eval.Var]bool, error) {
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	return vars, nil
}
