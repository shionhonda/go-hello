package intset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasAdd(t *testing.T) {
	var x IntSet
	assert.False(t, x.Has(1))
	x.Add(1)
	assert.True(t, x.Has(1))

	assert.False(t, x.Has(0))
	x.Add(0)
	assert.True(t, x.Has(0))

	assert.False(t, x.Has(64))
	x.Add(64)
	assert.Greater(t, len(x.words), 1)
	assert.True(t, x.Has(64))
}

func TestUnionWithString(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	y.Add(2)
	y.Add(3)
	y.Add(4)
	x.UnionWith(&y)
	assert.Equal(t, "{1 2 3 4}", x.String())
}

func TestLenRemoveClearCopy(t *testing.T) {
	var x IntSet
	assert.Equal(t, 0, x.Len())
	x.Remove(1)
	assert.Equal(t, 0, x.Len())
	x.Add(1)
	assert.Equal(t, 1, x.Len())
	x.Add(2)
	assert.Equal(t, 2, x.Len())
	x.Add(3)
	y := x.Copy()
	assert.Equal(t, 3, x.Len())
	assert.Equal(t, 3, y.Len())
	x.Remove(3)
	assert.Equal(t, 2, x.Len())
	x.Clear()
	assert.Equal(t, 0, x.Len())
	assert.Equal(t, 3, y.Len())
}
