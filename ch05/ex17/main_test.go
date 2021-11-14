package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	var ns []*html.Node
	f, _ := os.Open("fixture/go.html")
	defer f.Close()
	doc, _ := html.Parse(f)

	ns = ElementsByTagName(doc, "nav")
	assert.Len(t, ns, 1)

	ns = ElementsByTagName(doc, "h1", "div")
	assert.Len(t, ns, 15)

	ns = ElementsByTagName(doc, "thisshouldnotexist")
	assert.Len(t, ns, 0)
}
