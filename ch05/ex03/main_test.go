package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestFindTexts(t *testing.T) {
	f, _ := os.Open("fixture/go.html")
	defer f.Close()
	doc, _ := html.Parse(f)
	texts := findTexts(nil, doc)
	assert.Len(t, texts, 137)
}
