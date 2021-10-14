// main generates a PNG plot of Mandelbrot set
package main

import (
	"fmt"
	"go-hello/ch03/ex08/types"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 512, 512
)

func main() {
	var img *image.RGBA
	if os.Args[1] == "bigfloat" {
		img = imageBigFloat()
	} else if os.Args[1] == "bigrat" {
		img = imageBigRat()
	} else if os.Args[1] == "float32" {
		img = imageBigRat()
	} else {
		img = imageBigRat()
	}
	out, _ := os.Create(fmt.Sprintf("./%s.png", os.Args[1]))
	defer out.Close()
	png.Encode(out, img)

}

func imageBigFloat() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		yy := float64(py)/height*(ymax-ymin) + ymin
		y := big.NewFloat(yy)
		for px := 0; px < width; px++ {
			xx := float64(px)/width*(xmax-xmin) + xmin
			x := big.NewFloat(xx)
			z := &types.ComplexFloat{Real: x, Imag: y}
			img.Set(px, py, z.Mandelbrot())
		}
		fmt.Println(py)
	}
	return img
}

func imageBigRat() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := big.NewRat(int64(py)*(ymax-ymin), height)
		y.Add(y, big.NewRat(ymin, 1))
		for px := 0; px < width; px++ {
			x := big.NewRat(int64(px)*(xmax-xmin), width)
			x.Add(x, big.NewRat(xmin, 1))
			z := &types.ComplexRat{Real: x, Imag: y}
			img.Set(px, py, z.Mandelbrot())
		}
		fmt.Println(py)
	}
	return img
}

func imageFloat64() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot128(z))
		}
	}
	return img
}

func imageFloat32() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot64(z))
		}
	}
	return img
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 10

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - 10*n
			g := 170 - 15*n
			b := 85 - 20*n
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrot64(z complex64) color.Color {
	const iterations = 10

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if math.Sqrt(math.Pow(float64(real(v)), 2)+math.Pow(float64(imag(v)), 2)) > 2 {
			r := 255 - 10*n
			g := 170 - 15*n
			b := 85 - 20*n
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
