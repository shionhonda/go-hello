// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strings"
)

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
	// String returns a string representation of this Expr.
	String() string
}

//!+ast

// A Var identifies a variable, e.g., x.
type Var string

func (v Var) String() string {
	return string(v)
}

// A literal is a numeric constant, e.g., 3.141.
type literal float64

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x.String())
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.op, b.y.String())
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) String() string {
	var b strings.Builder
	b.WriteString(c.fn)
	b.WriteString("(")
	for i, arg := range c.args {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}
	b.WriteString(")")
	return b.String()
}
