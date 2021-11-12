package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopoSort(t *testing.T) {
	order := map[string]int{}
	for i, course := range topoSort(prereqs) {
		order[course] = i
	}
	for post, pres := range prereqs {
		for _, pre := range pres {
			assert.Greater(t, order[post], order[pre])
		}
	}
}
