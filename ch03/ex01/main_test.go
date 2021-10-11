package main

import (
	"math"
	"testing"
)

func TestCorner(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title  string
		input  []int
		output int
	}{
		{
			title:  "2×3の答えが6になる",
			input:  []int{2, 3},
			output: 6,
		},
		{
			title:  "3×5の答えが15になる",
			input:  []int{3, 5},
			output: 15,
		},
	} {
		t.Run("Square:"+td.title, func(t *testing.T) {
			_, _, err := corner(td.input[0], td.input[1])
			asserts.Equal(td.output, result)
		})
	}
	x, y, err := corner(0, 0, pi)

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
