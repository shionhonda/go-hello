package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	var tr *tree
	for i := 0; i < 10; i++ {
		tr = add(tr, i)
	}
	assert.Equal(t, tr.String(), "[0 1 2 3 4 5 6 7 8 9]")
}

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	for i := 0; i < len(data)-1; i++ {
		assert.GreaterOrEqual(t, data[i+1], data[i])
	}
}
