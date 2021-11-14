package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrJoin(t *testing.T) {
	var tests = []struct {
		sep   string
		elems []string
		want  string
	}{
		{"", []string{"foo", "bar"}, "foobar"},
		{" ", []string{"foo", "bar", "aa"}, "foo bar aa"},
		{" ", []string{}, ""},
	}
	for _, test := range tests {
		got := strJoin(test.sep, test.elems...)
		assert.Equal(t, test.want, got)
	}
}
