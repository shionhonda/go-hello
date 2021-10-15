package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountChar(t *testing.T) {
	f, _ := os.Open("fixture/sample.txt")
	defer f.Close()

	reader := bufio.NewReader(f)
	counter := countChar(reader)
	want := map[string]int{"letter": 11, "number": 9, "space": 18}
	assert.Len(t, counter, len(want))
	for k, v := range counter {
		assert.Equal(t, want[k], v)
	}
}
