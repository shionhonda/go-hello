package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearestRoot(t *testing.T) {
	var idx int
	var z complex128

	z = complex(1.1, 0.)
	idx = nearestRoot(z)
	assert.Equal(t, 0, idx)

	z = complex(1/2., 100.)
	idx = nearestRoot(z)
	assert.Equal(t, 1, idx)

	z = complex(1/2., -1)
	idx = nearestRoot(z)
	assert.Equal(t, 2, idx)
}
