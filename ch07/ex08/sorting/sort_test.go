package sorting

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	mc := NewMultiColumn(Tracks, [3]string{"Title", "Artist", "Album"})
	assert.EqualValues(t, [3]string{"Title", "Artist", "Album"}, mc.keys)
	sort.Sort(mc)
	assert.Equal(t, "Go", mc.t[0].Title)
	assert.Equal(t, "Asia", mc.t[0].Artist)
	assert.Equal(t, "Go", mc.t[1].Title)
	assert.Equal(t, "Delilah", mc.t[1].Artist)
	assert.Equal(t, "Go", mc.t[2].Title)
	assert.Equal(t, "Moby", mc.t[2].Artist)
	assert.Equal(t, "Go Ahead", mc.t[3].Title)
	assert.Equal(t, "Alicia Keys", mc.t[3].Artist)
	assert.Equal(t, "Ready 2 Go", mc.t[4].Title)
	assert.Equal(t, "Martin Solveig", mc.t[4].Artist)
	assert.Equal(t, "We Got The Beat", mc.t[5].Title)
	assert.Equal(t, "Go-Go's", mc.t[5].Artist)

	mc.SetPrimary("Year")
	assert.EqualValues(t, [3]string{"Year", "Title", "Artist"}, mc.keys)
	sort.Sort(mc)
	assert.Equal(t, "We Got The Beat", mc.t[0].Title)
	assert.Equal(t, "Go-Go's", mc.t[0].Artist)
	assert.Equal(t, "Go", mc.t[1].Title)
	assert.Equal(t, "Asia", mc.t[1].Artist)
	assert.Equal(t, "Go", mc.t[2].Title)
	assert.Equal(t, "Moby", mc.t[2].Artist)
	assert.Equal(t, "Go Ahead", mc.t[3].Title)
	assert.Equal(t, "Alicia Keys", mc.t[3].Artist)
	assert.Equal(t, "Ready 2 Go", mc.t[4].Title)
	assert.Equal(t, "Martin Solveig", mc.t[4].Artist)
	assert.Equal(t, "Go", mc.t[5].Title)
	assert.Equal(t, "Delilah", mc.t[5].Artist)
}
