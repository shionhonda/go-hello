// main generates a PNG plot of Newton's fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
)

var (
	z0 = complex(1, 0)
	z1 = complex(-1./2., math.Sqrt(3.)/2.)
	z2 = complex(-1./2., -math.Sqrt(3.)/2.)
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const (
		iterations = 20
	)
	v := z*z*z - 1
	for n := uint8(0); n < iterations; n++ {
		// Newton's method
		z = z - (v)/(3*z*z)
		v = z*z*z - 1
		if cmplx.Abs(v) < 0.01 {
			intensity := uint8(float64(iterations-n) / float64(iterations) * 255)
			idx := nearestRoot(z)
			if idx == 0 {
				return color.RGBA{intensity, 0, 0, 255}
			} else if idx == 1 {
				return color.RGBA{0, intensity, 0, 255}
			} else {
				return color.RGBA{0, 0, intensity, 255}
			}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func nearestRoot(z complex128) int {
	dists := []float64{cmplx.Abs(z - z0), cmplx.Abs(z - z1), cmplx.Abs(z - z2)}
	return argMin(dists)
}

func argMin(target []float64) int {
	var (
		index int
		base  float64
	)
	for i, d := range target {
		if i == 0 {
			index = i
			base = d
		} else {
			if d < base {
				index = i
				base = d
			}
		}

	}
	return index
}
