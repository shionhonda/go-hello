package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s = rotate(s, 0)
	e := []int{1, 2, 3, 4, 5}
	assert.EqualValues(t, e, s)

	s = rotate(s, 2)
	e = []int{3, 4, 5, 1, 2}
	assert.EqualValues(t, e, s)

	s = rotate(s, 6)
	e = []int{4, 5, 1, 2, 3}
	assert.EqualValues(t, e, s)
}
