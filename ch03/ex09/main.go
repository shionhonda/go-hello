package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// e.g., http://localhost:8000/?x=1&y=0.1&mag=2.0
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		x, y, mag := parseRequest(r)
		mandelbrot(w, x, y, mag)
	}
	http.HandleFunc("/", handler)
	log.Println("Server is running at http://localhost:8000/")
	log.Println("Assign parameters like: http://localhost:8000/?x=-1&y=0.1&mag=2.0")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseRequest(r *http.Request) (float64, float64, float64) {
	var (
		x   float64 = 0
		y   float64 = 0
		mag float64 = 1.0
	)

	for k, vs := range r.URL.Query() {
		if k == "x" {
			d, err := strconv.ParseFloat(vs[0], 64)
			if err == nil {
				x = d
			}
		}
		if k == "y" {
			d, err := strconv.ParseFloat(vs[0], 64)
			if err == nil {
				y = d
			}
		}
		if k == "mag" {
			d, err := strconv.ParseFloat(vs[0], 64)
			if err == nil && d > 0 {
				mag = d
			}
		}
	}
	return x, y, mag
}

func mandelbrot(out io.Writer, cx float64, cy float64, mag float64) {
	const (
		width, height = 1024, 1024
	)
	var (
		xmin = cx - 2/mag
		xmax = cx + 2/mag
		ymin = cy - 2/mag
		ymax = cy + 2/mag
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, getColor(z))
		}
	}

	png.Encode(out, img)
}

func getColor(z complex128) color.Color {
	const iterations = 200

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
