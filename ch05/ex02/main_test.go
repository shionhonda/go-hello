package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestCountElement(t *testing.T) {
	f, _ := os.Open("fixture/go.html")
	defer f.Close()
	doc, _ := html.Parse(f)
	m := countElements(map[string]int{}, doc)
	assert.Len(t, m, 30)
	assert.Equal(t, 17, m["a"])
	assert.Equal(t, 1, m["p"])
	assert.Equal(t, 14, m["div"])
	assert.Equal(t, 0, m["spam"])
}
