package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"hello", false},
		{"detereted", true},
		{"vv", true},
		{"aka", true},
		{"ab", false},
		{"", true},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, IsPalindrome(newMyStrings(test.input)))
	}
}
