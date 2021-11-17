package main

import (
	"fmt"
	"go-hello/ch07/ex01/counter"
	"io"
)

type countingWriter struct {
	w io.Writer
	n int64
}

func main() {
	var lc counter.LineCounter
	cw, naddr := CountingWriter(&lc)
	fmt.Println(*naddr)
	cw.Write([]byte("hello"))
	fmt.Println(*naddr)
	cw.Write([]byte("world"))
	fmt.Println(*naddr)
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.n += int64(n)
	return n, err
}

// CountingWriter is a wrapper for io.Writer that contains the number of bytes written.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w: w}
	return cw, &cw.n
}
