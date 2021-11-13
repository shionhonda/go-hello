package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	f, _ := os.Open("fixture/go.html")
	defer f.Close()
	doc, _ := html.Parse(f)
	links := visit(nil, doc)
	assert.Len(t, links, 22)
}
