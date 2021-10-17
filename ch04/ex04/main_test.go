package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	rotate(s, 0)
	e := []int{1, 2, 3, 4, 5, 6}
	assert.EqualValues(t, e, s)

	rotate(s, 2)
	e = []int{3, 4, 5, 6, 1, 2}
	assert.EqualValues(t, e, s)

	rotate(s, 5)
	e = []int{2, 3, 4, 5, 6, 1}
	assert.EqualValues(t, e, s)

	rotate(s, 6)
	e = []int{2, 3, 4, 5, 6, 1}
	assert.EqualValues(t, e, s)

	rotate(s, 7)
	e = []int{3, 4, 5, 6, 1, 2}
	assert.EqualValues(t, e, s)
}

func TestJuggle(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	juggle(s, 0, 2)
	e := []int{3, 4, 5, 1, 2}
	assert.EqualValues(t, e, s)

	juggle(s, 0, 1)
	e = []int{4, 5, 1, 2, 3}
	assert.EqualValues(t, e, s)

	s = []int{1, 2, 3, 4, 5, 6}
	juggle(s, 1, 3)
	e = []int{1, 5, 3, 4, 2, 6}
	assert.EqualValues(t, e, s)
}

func TestGCD(t *testing.T) {
	var a, b, e int

	a = 7
	b = 3
	e = 1
	assert.Equal(t, e, gcd(a, b))
	assert.Equal(t, e, gcd(b, a))

	a = 24
	b = 36
	e = 12
	assert.Equal(t, e, gcd(a, b))
	assert.Equal(t, e, gcd(b, a))
}
