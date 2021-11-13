package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopoSort(t *testing.T) {
	var tests = []struct {
		name  string
		input map[string][]string
		msg   string
	}{
		{"valid", map[string][]string{
			"algorithms": {"data structures"},
			"calculus":   {"linear algebra"},
			"compilers": {
				"data structures",
				"formal languages",
				"computer organization",
			},
			"data structures":       {"discrete math"},
			"databases":             {"data structures"},
			"discrete math":         {"intro to programming"},
			"formal languages":      {"discrete math"},
			"networks":              {"operating systems"},
			"operating systems":     {"data structures", "computer organization"},
			"programming languages": {"data structures", "computer organization"},
		}, ""},
		{"cyclic", map[string][]string{
			"calculus":       {"linear algebra"},
			"linear algebra": {"calculus"},
		}, "cycle detected"},
		{"cyclic2", map[string][]string{
			"algorithms": {"data structures"},
			"calculus":   {"linear algebra"},
			"compilers": {
				"data structures",
				"formal languages",
				"computer organization",
			},
			"data structures":       {"discrete math"},
			"databases":             {"data structures"},
			"discrete math":         {"intro to programming"},
			"formal languages":      {"discrete math"},
			"networks":              {"operating systems"},
			"operating systems":     {"data structures", "computer organization"},
			"programming languages": {"data structures", "computer organization"},
			"linear algebra":        {"calculus"},
		}, "cycle detected"},
	}
	for _, test := range tests {
		courses, err := topoSort(test.input)
		if test.msg != "" {
			assert.EqualError(t, err, test.msg)
		} else {
			assert.NoError(t, err)
			order := map[string]int{}
			for i, course := range courses {
				order[course] = i
			}
			for post, pres := range test.input {
				for _, pre := range pres {
					assert.Greater(t, order[post], order[pre])
				}
			}
		}
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		list    []string
		x       string
		contain bool
	}{
		{[]string{"aa", "bb", "cc"}, "aa", true},
		{[]string{"aa", "bb", "cc"}, "dd", false},
	}
	for _, test := range tests {
		contain := contains(test.list, test.x)
		assert.Equal(t, test.contain, contain)
	}
}
