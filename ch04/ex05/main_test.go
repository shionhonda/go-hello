package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	var s, e []string

	s = []string{"1", "2", "2", "3", "4", "5"}
	s = removeNextDups(s)
	e = []string{"1", "2", "3", "4", "5"}
	assert.EqualValues(t, e, s)

	s = []string{"2", "2", "2", "2", "2", "2", "2", "2", "2", "2"}
	s = removeNextDups(s)
	e = []string{"2"}
	assert.EqualValues(t, e, s)

	s = []string{"1", "3", "5", "8", "2", "4"}
	s = removeNextDups(s)
	e = []string{"1", "3", "5", "8", "2", "4"}
	assert.EqualValues(t, e, s)

	s = []string{"1", "3", "3", "8", "2", "4", "4", "1", "5"}
	s = removeNextDups(s)
	e = []string{"1", "3", "8", "2", "4", "1", "5"}
	assert.EqualValues(t, e, s)

	s = []string{"1", "3", "5", "8", "2", "2"}
	s = removeNextDups(s)
	e = []string{"1", "3", "5", "8", "2"}
	assert.EqualValues(t, e, s)

}
