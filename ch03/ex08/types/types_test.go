package types

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexFloat(t *testing.T) {
	var f float64
	delta := 0.001
	z0 := &ComplexFloat{big.NewFloat(3.), big.NewFloat(2.)}
	z1 := &ComplexFloat{big.NewFloat(-2.), big.NewFloat(-4.)}

	z0.mul(z1) // 2 - 16i
	f, _ = z0.Real.Float64()
	assert.InDelta(t, 2., f, delta)
	f, _ = z0.Imag.Float64()
	assert.InDelta(t, -16., f, delta)

	z0.add(z1) // - 20i
	f, _ = z0.Real.Float64()
	assert.InDelta(t, 0., f, delta)
	f, _ = z0.Imag.Float64()
	assert.InDelta(t, -20., f, delta)

	abs := z0.abs() // 20
	f, _ = abs.Float64()
	assert.InDelta(t, 20, f, delta)

	z := &ComplexFloat{big.NewFloat(0.3), big.NewFloat(0.6)}
	c := z.Mandelbrot()
	r, g, b, a := c.RGBA()
	assert.Equal(t, uint8(115), uint8(r))
	assert.Equal(t, uint8(216), uint8(g))
	assert.Equal(t, uint8(61), uint8(b))
	assert.Equal(t, uint8(255), uint8(a))
}

func TestComplexRat(t *testing.T) {
	var f float64
	delta := 0.001
	z0 := &ComplexRat{big.NewRat(3, 1), big.NewRat(2, 1)}
	z1 := &ComplexRat{big.NewRat(-2, 1), big.NewRat(-4, 1)}

	z0.mul(z1) // 2 - 16i
	f, _ = z0.Real.Float64()
	assert.InDelta(t, 2., f, delta)
	f, _ = z0.Imag.Float64()
	assert.InDelta(t, -16., f, delta)

	z0.add(z1) // - 20i
	f, _ = z0.Real.Float64()
	assert.InDelta(t, 0., f, delta)
	f, _ = z0.Imag.Float64()
	assert.InDelta(t, -20., f, delta)

	abs := z0.abs() // 20
	assert.InDelta(t, 20, abs, delta)

	z := &ComplexRat{big.NewRat(3, 10), big.NewRat(6, 10)}
	c := z.Mandelbrot()
	r, g, b, a := c.RGBA()
	assert.Equal(t, uint8(115), uint8(r))
	assert.Equal(t, uint8(216), uint8(g))
	assert.Equal(t, uint8(61), uint8(b))
	assert.Equal(t, uint8(255), uint8(a))
}
