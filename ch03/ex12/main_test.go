package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	title    string
	s0       string
	s1       string
	expected bool
}

func TestIsAnagram(t *testing.T) {
	testCases := []testCase{
		{"abc", "abc", "cba", true},
		{"abC", "abC", "cba", false},
		{"abcc", "abcc", "cbca", true},
		{"abcc", "abcc", "cbaa", false},
		{"abcc", "c", "cbaa", false},
	}
	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			b := isAnagram(tc.s0, tc.s1)
			assert.Equal(t, tc.expected, b)
		})
	}
}
