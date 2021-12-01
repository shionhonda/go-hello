package main

import (
	"fmt"
	"go-hello/ch07/ex14/eval"
	"html/template"
	"log"
	"net/http"
)

type Result struct {
	Value float64
}

func main() {
	http.HandleFunc("/", calculate)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		calculator.Execute(w, nil)
	} else {
		r.ParseForm()
		exprs := r.Form["expression"]
		if len(exprs) == 0 {
			fmt.Fprintf(w, "no expression")
			calculator.Execute(w, nil)
		} else {
			expr, err := eval.Parse(exprs[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest) // 400
				fmt.Fprintf(w, "parse error: %v\n", err)
				return
			}
			value := expr.Eval(eval.Env{})
			fmt.Printf("Expression: %s\n", expr)
			fmt.Printf("Value: %f\n", value)
			calculator.Execute(w, Result{value})
		}

	}

}

var calculator = template.Must(template.New("calculator").Parse(`
<html>
<head>
<title>Calculator</title>
</head>
<body>
<h1>Calculator</h1>
<form action="/" method="post">
<input type="text" name="expression", value={{.Value}}>
<input type="submit" value="Enter">
</form>
</body>
</html>
`))
