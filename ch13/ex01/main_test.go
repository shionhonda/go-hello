package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	x     interface{}
	y     interface{}
	equal bool
}

func TestNearlyEqual(t *testing.T) {
	testCases := []testCase{
		{1.0, 1.0, true},
		{1.0, 1.0 + 9.9e-10, true},
		{1.0, 1.0 - 9.9e-10, true},
		{4, 4 + 9.9e-10, true},
		{-4, -4 - 9.9e-10, true},
		{1.0, 1.0 - 1.1e-9, false},
		{4, 4 + 1.1e-9, false},
	}
	for _, tc := range testCases {
		if tc.equal {
			assert.True(t, NearlyEqual(tc.x, tc.y))
		} else {
			assert.False(t, NearlyEqual(tc.x, tc.y))
		}
	}
}
