// mandelbrot geenerates a PNG plot of Mandelbrot set
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
)

func main() {
	rec := createExtendedRectangle()

	// Average surrounding pixels
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var r, g, b uint8
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					r += rec[2*py+i][2*px+j][0] / 4
					g += rec[2*py+i][2*px+j][1] / 4
					b += rec[2*py+i][2*px+j][2] / 4
				}
			}
			img.Set(px, py, color.RGBA{r, g, b, 255})
		}
	}

	png.Encode(os.Stdout, img)
}

func createExtendedRectangle() [2 * height][2 * width][3]uint8 {
	rec := [2 * height][2 * width][3]uint8{}
	for py := 0; py < 2*height; py++ {
		y := float64(py)/2/height*(ymax-ymin) + ymin
		for px := 0; px < 2*width; px++ {
			x := float64(px)/2/width*(xmax-xmin) + xmin
			z := complex(x, y)
			rec[py][px] = mandelbrot(z)
		}
	}
	return rec
}

func mandelbrot(z complex128) [3]uint8 {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - 10*n
			g := 170 - 15*n
			b := 85 - 20*n
			return [3]uint8{r, g, b}
		}
	}
	return [3]uint8{0, 0, 0}
}
