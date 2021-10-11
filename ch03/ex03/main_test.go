package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	title string
	f     func(float64, float64) float64
	isErr bool
}

func TestCorner(t *testing.T) {
	testCases := []testCase{
		{"No error", pi, false},
		{"Positive infinity", posInf, true},
		{"Negative infinity", negInf, true},
		{"NaN", nan, true},
	}
	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			_, _, _, err := corner(0, 0, tc.f)
			if tc.isErr {
				assert.NotEqual(t, nil, err)
			} else {
				assert.Equal(t, nil, err)
			}
		})
	}
}

func TestColorCode(t *testing.T) {
	c := colorCode(maxh)
	assert.Equal(t, c, "#ff0000")

	c = colorCode(minh)
	assert.Equal(t, c, "#0000ff")
}

func pi(x, y float64) float64 {
	return math.Pi
}

func posInf(x, y float64) float64 {
	return math.Inf(1)
}

func negInf(x, y float64) float64 {
	return math.Inf(-1)
}

func nan(x, y float64) float64 {
	return math.NaN()
}
