// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	p := printer{width: margin}
	if err := encode(&p, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return p.Bytes(), nil
}

const margin = 80

type token struct {
	kind rune // one of "s () {} []" (string, blank, start, end)
	str  string
	size int
}

type printer struct {
	tokens []*token // FIFO buffer
	stack  []*token // stack of open ' ' and '(' tokens
	rtotal int      // total number of spaces needed to print stream

	bytes.Buffer
	indents []int
	width   int // remaining space
}

func (p *printer) string(str string) {
	tok := &token{kind: 's', str: str, size: len(str)}
	if len(p.stack) == 0 {
		p.print(tok)
	} else {
		p.tokens = append(p.tokens, tok)
		p.rtotal += len(str)
	}
}
func (p *printer) pop() (top *token) {
	last := len(p.stack) - 1
	top, p.stack = p.stack[last], p.stack[:last]
	return
}
func (p *printer) begin(r rune) {
	if len(p.stack) == 0 {
		p.rtotal = 1
	}
	t := &token{kind: r, size: -p.rtotal}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t) // push
	p.string(string(r))
}
func (p *printer) end(r rune) {
	p.string(string(r))
	p.tokens = append(p.tokens, &token{kind: r})
	x := p.pop()
	x.size += p.rtotal
	if x.kind == ' ' {
		p.pop().size += p.rtotal
	}
	if len(p.stack) == 0 {
		for _, tok := range p.tokens {
			p.print(tok)
		}
		p.tokens = nil
	}
}

func (p *printer) space() {
	last := len(p.stack) - 1
	x := p.stack[last]
	if x.kind == ' ' {
		x.size += p.rtotal
		p.stack = p.stack[:last] // pop
	}
	t := &token{kind: ' ', size: -p.rtotal}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t)
	p.rtotal++
}
func (p *printer) print(t *token) {
	switch t.kind {
	case 's':
		p.WriteString(t.str)
		p.width -= len(t.str)
	case '[', '{':
		p.indents = append(p.indents, p.width)
	case ']', '}':
		p.indents = p.indents[:len(p.indents)-1] // pop
	case ' ':
		if t.size > p.width {
			p.width = p.indents[len(p.indents)-1] - 1
			fmt.Fprintf(&p.Buffer, "\n%*s", margin-p.width, "")
		} else {
			p.WriteByte(' ')
			p.width--
		}
	}
}
func (p *printer) stringf(format string, args ...interface{}) {
	p.string(fmt.Sprintf(format, args...))
}

func encode(p *printer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		p.string("null")

	case reflect.Bool:
		if v.Bool() {
			p.string("true")
		} else {
			p.string("false")
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		p.stringf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p.stringf("%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		p.stringf("%f", v.Float())

	case reflect.String:
		p.stringf("%q", v.String())

	case reflect.Array, reflect.Slice: // (value ...)
		p.begin('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				p.space()
			}
			if err := encode(p, v.Index(i)); err != nil {
				return err
			}
			if i < v.Len()-1 {
				p.string(", ")
			}
		}
		p.end(']')

	case reflect.Struct: // ((name value ...)
		p.begin('{')
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).IsZero() {
				continue // do nothing
			}
			if i > 0 {
				p.string(", ")
				p.space()
			}
			p.stringf("%q: ", v.Type().Field(i).Name)
			p.space()
			if err := encode(p, v.Field(i)); err != nil {
				return err
			}

		}
		p.end('}')

	case reflect.Map: // ((key value ...)
		p.begin('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				p.space()
			}
			if err := encode(p, key); err != nil {
				return err
			}
			p.string(": ")
			p.space()
			if err := encode(p, v.MapIndex(key)); err != nil {
				return err
			}
			if i < v.Len()-1 {
				p.string(", ")
			}
		}
		p.end('}')

	case reflect.Ptr:
		return encode(p, v.Elem())

	default: // chan, func, complex, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}