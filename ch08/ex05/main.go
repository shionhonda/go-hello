// main generates a PNG plot of Mandelbrot set
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

type data struct {
	px    int
	py    int
	color color.Color
}

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	ch := make(chan data, 1024*1024)

	go func() {
		for py := 0; py < height; py++ {
			for px := 0; px < width; px++ {
				mandelbrot(px, py, ch)
			}
		}
		close(ch)
	}()

	for d := range ch {
		img.Set(d.px, d.py, d.color)
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(px, py int, out chan<- data) {
	y := float64(py)/height*(ymax-ymin) + ymin
	x := float64(px)/width*(xmax-xmin) + xmin
	z := complex(x, y)

	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - 10*n
			g := 170 - 15*n
			b := 85 - 20*n
			out <- data{px, py, color.RGBA{r, g, b, 255}}
			return
		}
	}
	out <- data{px, py, color.RGBA{0, 0, 0, 255}}
}
