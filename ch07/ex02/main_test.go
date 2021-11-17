package main

import (
	"go-hello/ch07/ex01/counter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingWriter(t *testing.T) {
	var lc counter.LineCounter
	cw, naddr := CountingWriter(&lc)
	assert.Equal(t, int64(0), *naddr)
	cw.Write([]byte("hello"))
	assert.Equal(t, int64(5), *naddr)
	cw.Write([]byte("world"))
	assert.Equal(t, int64(10), *naddr)
}
