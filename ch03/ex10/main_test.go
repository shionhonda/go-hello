package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	title    string
	num      string
	expected string
}

func TestComma(t *testing.T) {
	testCases := []testCase{
		{"Two commas", "123456789", "123,456,789"},
		{"One comma", "1234", "1,234"},
		{"No comma", "123", "123"},
	}
	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			s := comma(tc.num)
			assert.Equal(t, tc.expected, s)
		})
	}
}
