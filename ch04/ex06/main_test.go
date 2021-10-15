package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	var s, r, e string

	s = "1  2 3  4"
	r = string(removeDupSpace([]byte(s)))
	e = "1 2 3 4"
	assert.Equal(t, e, r)

	// NOTE: This does not pass
	s = "1  2 3      4"
	r = string(removeDupSpace([]byte(s)))
	e = "1 2 3 4"
	assert.Equal(t, e, r)

}
