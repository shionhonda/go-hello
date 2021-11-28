package eval

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		{"min(x, y, 3)", Env{"x": 9, "y": -1}, "-1"},
		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		fmt.Println(expr)
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		assert.Equal(t, test.want, got)

		// Check again
		exprStr := fmt.Sprintf("%s", expr)
		expr2, err := Parse(exprStr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		assert.Equal(t, expr, expr2)
	}
}

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}
