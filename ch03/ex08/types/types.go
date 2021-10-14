package types

import (
	"image/color"
	"math"
	"math/big"
)

type ComplexFloat struct {
	Real *big.Float
	Imag *big.Float
}

type ComplexRat struct {
	Real *big.Rat
	Imag *big.Rat
}

func (z *ComplexFloat) Mandelbrot() color.Color {
	const iterations = 10

	v := &ComplexFloat{big.NewFloat(0), big.NewFloat(0)}
	for n := uint8(0); n < iterations; n++ {
		v.mul(v)
		v.add(z)
		abs := v.abs()
		if abs.Cmp(big.NewFloat(2.0)) == 1 {
			r := 255 - 10*n
			g := 170 - 15*n
			b := 85 - 20*n
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func (z0 *ComplexFloat) mul(z1 *ComplexFloat) {
	var r, ra, rb, i, ia, ib big.Float
	ra.Mul(z0.Real, z1.Real)
	rb.Mul(z0.Imag, z1.Imag)
	r.Sub(&ra, &rb)
	ia.Mul(z0.Real, z1.Imag)
	ib.Mul(z0.Imag, z1.Real)
	i.Add(&ia, &ib)
	z0.Real = &r
	z0.Imag = &i
}

func (z0 *ComplexFloat) add(z1 *ComplexFloat) {
	var r, i big.Float
	r.Add(z0.Real, z1.Real)
	i.Add(z0.Imag, z1.Imag)
	z0.Real = &r
	z0.Imag = &i
}

func (z *ComplexFloat) abs() big.Float {
	var r, i, c, ret big.Float
	r.Mul(z.Real, z.Real)
	i.Mul(z.Imag, z.Imag)
	c.Add(&r, &i)
	ret.Sqrt(&c)
	return ret
}

func (z *ComplexRat) Mandelbrot() color.Color {
	const iterations = 10

	v := &ComplexRat{big.NewRat(0, 1), big.NewRat(0, 1)}
	for n := uint8(0); n < iterations; n++ {
		v.mul(v)
		v.add(z)
		abs := v.abs()
		if abs > 2 {
			r := 255 - 10*n
			g := 170 - 15*n
			b := 85 - 20*n
			return color.RGBA{r, g, b, 255}
		} else if abs < 0.001 {
			return color.RGBA{0, 0, 0, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func (z0 *ComplexRat) mul(z1 *ComplexRat) {
	var r, ra, rb, i, ia, ib big.Rat
	ra.Mul(z0.Real, z1.Real)
	rb.Mul(z0.Imag, z1.Imag)
	r.Sub(&ra, &rb)
	ia.Mul(z0.Real, z1.Imag)
	ib.Mul(z0.Imag, z1.Real)
	i.Add(&ia, &ib)
	z0.Real = &r
	z0.Imag = &i
}

func (z0 *ComplexRat) add(z1 *ComplexRat) {
	var r, i big.Rat
	r.Add(z0.Real, z1.Real)
	i.Add(z0.Imag, z1.Imag)
	z0.Real = &r
	z0.Imag = &i
}

func (z *ComplexRat) abs() float64 {
	var r, i, c big.Rat
	r.Mul(z.Real, z.Real)
	i.Mul(z.Imag, z.Imag)
	c.Add(&r, &i)
	f, _ := c.Float64()
	return math.Sqrt(f)
}
