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

	sc := bufio.NewScanner(f)
	counter := countWords(sc)
	assert.Equal(t, 14, counter["the"])
	assert.Equal(t, 9, counter["i"])
}
