package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitedReader(t *testing.T) {
	r := NewReader("Hello, Reader!")
	lr := LimitReader(r, 5)
	b := make([]byte, 8)

	n, err := lr.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "Hello", string(b[:n]))

	n, err = lr.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, ", Rea", string(b[:n]))

	n, err = lr.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, "der!", string(b[:n]))

	n, err = lr.Read(b)
	assert.ErrorIs(t, err, io.EOF)
	assert.Equal(t, 0, n)
}
