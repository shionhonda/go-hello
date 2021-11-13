package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
	f, _ := os.Open("fixture/go.html")
	defer f.Close()
	doc, _ := html.Parse(f)
	words, images := countWordsAndImages(doc)
	assert.Equal(t, 478, words)
	assert.Equal(t, 3, images)
}
