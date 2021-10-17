package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	var s, r, e string

	s = "a  b c  e"
	r = string(removeDupSpace([]byte(s)))
	e = "a b c e"
	assert.Equal(t, e, r)

	s = "a  b c     e"
	r = string(removeDupSpace([]byte(s)))
	e = "a b c e"
	assert.Equal(t, e, r)

	s = "1  2 3      あ"
	r = string(removeDupSpace([]byte(s)))
	e = "1 2 3 あ"
	assert.Equal(t, e, r)

	s = "1  2 3 \t　   あ"
	r = string(removeDupSpace([]byte(s)))
	e = "1 2 3 あ"
	assert.Equal(t, e, r)

	s = "a   \t"
	r = string(removeDupSpace([]byte(s)))
	e = "a "
	assert.Equal(t, e, r)

	s = "  a   \t"
	r = string(removeDupSpace([]byte(s)))
	e = " a "
	assert.Equal(t, e, r)

	s = " \t　   \n "
	r = string(removeDupSpace([]byte(s)))
	e = " "
	assert.Equal(t, e, r)

}
