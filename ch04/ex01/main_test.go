package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingDist(t *testing.T) {
	a := [32]byte{byte(255)}
	b := [32]byte{byte(3)}
	e := 6 // 11111111 v.s. 00000011
	diff := hammingDist(a, b)
	assert.Equal(t, e, diff)
}

func TestPopCount(t *testing.T) {
	var x uint64
	var y, e int

	x = 127
	e = 7
	y = popCount(x)
	assert.Equal(t, e, y)

	x = 1234567890
	e = 12
	y = popCount(x)
	assert.Equal(t, e, y)

	x = 0
	e = 0
	y = popCount(x)
	assert.Equal(t, e, y)

	x = 1
	e = 1
	y = popCount(x)
	assert.Equal(t, e, y)
}
