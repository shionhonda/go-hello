package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	s := [5]int{1, 2, 3, 4, 5}
	reverse(&s)
	e := [5]int{5, 4, 3, 2, 1}
	assert.EqualValues(t, e, s)
}
