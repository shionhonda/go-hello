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
		{"Two commas", "123456789.12", "123,456,789.12"},
		{"One comma", "-1234.5678", "-1,234.5678"},
		{"No comma", "-123.2", "-123.2"},
	}
	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			s := comma(tc.num)
			assert.Equal(t, tc.expected, s)
		})
	}
}
