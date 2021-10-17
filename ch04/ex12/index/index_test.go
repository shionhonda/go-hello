package index

import (
	"testing"

	"go-hello/ch04/ex12/mytypes"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	comic := LoadJSON("../fixture/571.json")
	assert.Equal(t, 571, comic.Num)
	assert.Equal(t, "Can't Sleep", comic.Title)

	m := *Index(&[]mytypes.Comic{*comic})
	_, ok := m["sleep"]
	assert.True(t, ok)

	_, ok = m["do"]
	assert.True(t, ok)

	_, ok = m["hoge"]
	assert.False(t, ok)
}

func TestSplit(t *testing.T) {
	ws := split("I have a pen, and she has an erasor!")
	expected := []string{"I", "have", "a", "pen", "and", "she", "has", "an", "erasor"}
	assert.EqualValues(t, expected, ws)
}
