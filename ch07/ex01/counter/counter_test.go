package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounter(t *testing.T) {
	var c WordCounter
	assert.Equal(t, 0, int(c))
	c.Write([]byte("goodbye 昨日"))
	assert.Equal(t, 2, int(c))
	c.Write([]byte("hello 今日"))
	assert.Equal(t, 4, int(c))
}

func TestLineCounter(t *testing.T) {
	var c LineCounter
	assert.Equal(t, 0, int(c))
	c.Write([]byte("goodbye 昨日"))
	assert.Equal(t, 1, int(c))
	c.Write([]byte("hello\n今日"))
	assert.Equal(t, 3, int(c))
}
