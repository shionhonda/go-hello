package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	var n *html.Node
	f, _ := os.Open("fixture/go.html")
	defer f.Close()
	doc, _ := html.Parse(f)

	n = ElementByID(doc, "nav")
	assert.Equal(t, "div", n.Data)

	n = ElementByID(doc, "page")
	assert.Equal(t, "main", n.Data)

	n = ElementByID(doc, "thisshouldnotexist")
	assert.Nil(t, n)
}
