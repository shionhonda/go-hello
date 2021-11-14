package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		vals []int
		want int
	}{
		{[]int{1, 3, 5, 7, 9}, 9},
		{[]int{-1, -3, -5, -7, -9}, -1},
		{[]int{-2, -4, 0, 4}, 4},
		{[]int{-1, -5, 0}, 0},
		{[]int{}, minInt},
	}
	for _, test := range tests {
		m := max(test.vals...)
		assert.Equal(t, test.want, m)
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		vals []int
		want int
	}{
		{[]int{1, 3, 5, 7, 9}, 1},
		{[]int{-1, 3, -5, 7, -9}, -9},
		{[]int{-2, -4, 0, 4}, -4},
		{[]int{1, 5, 0}, 0},
		{[]int{}, maxInt},
	}
	for _, test := range tests {
		m := min(test.vals...)
		assert.Equal(t, test.want, m)
	}
}

func TestSafeMax(t *testing.T) {
	var tests = []struct {
		vals  []int
		want  int
		isErr bool
	}{
		{[]int{1, 3, 5, 7, 9}, 9, false},
		{[]int{-1, -3, -5, -7, -9}, -1, false},
		{[]int{-2, -4, 0, 4}, 4, false},
		{[]int{-1, -5, 0}, 0, false},
		{[]int{}, 0, true},
	}
	for _, test := range tests {
		m, err := safeMax(test.vals...)
		if test.isErr {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.want, m)
		}
	}
}

func TestSafeMin(t *testing.T) {
	var tests = []struct {
		vals  []int
		want  int
		isErr bool
	}{
		{[]int{1, 3, 5, 7, 9}, 1, false},
		{[]int{-1, 3, -5, 7, -9}, -9, false},
		{[]int{-2, -4, 0, 4}, -4, false},
		{[]int{1, 5, 0}, 0, false},
		{[]int{}, 0, true},
	}
	for _, test := range tests {
		m, err := safeMin(test.vals...)
		if test.isErr {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.want, m)
		}
	}
}
