package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	var tests = []struct {
		input string
		f     func(string) string
		want  string
	}{
		{"", strings.ToUpper, ""},
		{"abc", strings.ToUpper, "abc"},
		{"abc $foo def", strings.ToUpper, "abc FOO def"},
		{"abc $foo $bar $hogeeeee", strings.ToUpper, "abc FOO BAR HOGEEEEE"},
	}
	for _, test := range tests {
		if got := expand(test.input, test.f); got != test.want {
			assert.Equal(t, test.want, got)
		}
	}
}
