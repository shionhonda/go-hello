package main

import (
	"fmt"
	"io"
)

type LimitedReader struct {
	r io.Reader
	n int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func (l *LimitedReader) Read(p []byte) (int, error) {
	if len(p) > int(l.n) {
		p = p[:l.n]
	}
	return l.r.Read(p)
}

type Reader struct {
	s string
}

func NewReader(s string) *Reader {
	return &Reader{s}
}

func (r *Reader) Read(p []byte) (int, error) {
	if len(r.s) == 0 {
		return 0, io.EOF
	}
	n := copy(p, r.s)
	r.s = r.s[n:]
	return n, nil
}

func main() {
	lr := LimitReader(NewReader("abcdefghijklmnopqrstuvwxyz"), 5)
	buf := make([]byte, 5)
	for {
		n, err := lr.Read(buf)
		if err != nil {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("%s\n", string(buf[:n]))
	}

}
