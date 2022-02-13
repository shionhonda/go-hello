package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

// Token is one of: Symbol, String, Int, StartList, EndList
type Token interface{}
type Symbol string // foo
type String string // "hello"
type Int int       // 42
type StartList struct{}
type EndList struct{}

type Decoder struct {
	scan scanner.Scanner
	err  error
}

func NewDecoder(r io.Reader) *Decoder {
	var scan scanner.Scanner
	scan.Init(r)
	dec := Decoder{scan: scan}
	scan.Error = dec.setError
	return &dec
}

func (d *Decoder) setError(scan *scanner.Scanner, msg string) {
	d.err = fmt.Errorf("scanning: %s", msg)
}

func (d *Decoder) Token() (Token, error) {
	t := d.scan.Scan()
	if d.err != nil {
		return nil, d.err
	}
	switch t {
	case scanner.EOF:
		return nil, io.EOF
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		text := d.scan.TokenText()
		return Symbol(text), nil
	case scanner.String:
		text := d.scan.TokenText()
		// removing quatation marks
		return String(text[1 : len(text)-1]), nil
	case scanner.Int:
		i, _ := strconv.Atoi(d.scan.TokenText()) // NOTE: ignoring errors
		fmt.Println(i)
		return i, nil
	case '(':
		return StartList{}, nil
	case ')':
		return EndList{}, nil
	}
	panic(fmt.Sprintf("unexpected token %q", d.scan.TokenText()))
}

func (dec *Decoder) Decode(v interface{}) (err error) {
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.scan.Position, x)
		}
	}()
	w := reflect.ValueOf(v).Elem()
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch t := t.(type) {
		case Symbol:
			w.Set(reflect.Zero(w.Type()))
		case String:
			w.SetString(string(t))
		case Int:
			w.SetInt(int64(t))
		}
	}
	return nil
}

func Unmarshal(data []byte, out interface{}) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(out)
	return err
}
