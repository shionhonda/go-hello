package tempconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	var c Celsius
	var f Fahrenheit
	var k Kelvin

	c = Celsius(10)
	assert.Equal(t, "10˚C", c.String())

	c = Celsius(-10.4)
	assert.Equal(t, "-10.4˚C", c.String())

	f = Fahrenheit(100)
	assert.Equal(t, "100˚F", f.String())

	f = Fahrenheit(-20.333)
	assert.Equal(t, "-20.333˚F", f.String())

	k = Kelvin(300)
	assert.Equal(t, "300K", k.String())

	k = Kelvin(0.1)
	assert.Equal(t, "0.1K", k.String())
}
