package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	r := NewReader("Hello, Reader!")
	b := make([]byte, 8)
	n, err := r.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, 8, n)
	assert.Equal(t, "Hello, R", string(b))
}
