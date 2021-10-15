package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBytes(t *testing.T) {
	s := []byte("Hello")
	reverseBytes(s)
	e := []byte("olleH")
	assert.Equal(t, e, s)
}

func TestReverseUTF8(t *testing.T) {
	s := "Hello, 世界"
	s = reverseUTF8(s)
	e := "界世 ,olleH"
	assert.Equal(t, e, s)
}
