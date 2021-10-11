package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 60                  // number of grid cells
	xyrange       = 2.5                 // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	minh          = -1.031339
	maxh          = 7.451471
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var red = []uint8{255, 0, 0}
var blue = []uint8{0, 0, 255}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z, err := corner(i+1, j, sixhump)
			if err != nil {
				continue
			}
			bx, by, _, err := corner(i, j, sixhump)
			if err != nil {
				continue
			}
			cx, cy, _, err := corner(i, j+1, sixhump)
			if err != nil {
				continue
			}
			dx, dy, _, err := corner(i+1, j+1, sixhump)
			if err != nil {
				continue
			}
			c := colorCode(z)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}
	fmt.Println("</svg>")

}

func corner(i, j int, f func(float64, float64) float64) (float64, float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	// Return error if z in Inf or NaN
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0., 0., 0., fmt.Errorf("Invalid output: %v", z)
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func sixhump(x, y float64) float64 {
	x2 := math.Pow(x, 2)
	x4 := math.Pow(x, 4)
	y2 := math.Pow(y, 2)
	return ((4-2.1*x2+x4/3)*x2 + x*y + (-4+4*y2)*y2)
}

func colorCode(h float64) string {
	t := (h - minh) / (maxh - minh)
	r, _ := interpolate(red[0], blue[0], t)
	g := 0
	b, _ := interpolate(red[2], blue[2], t)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func interpolate(x, y uint8, t float64) (uint8, error) {
	if t < 0 || t > 1 {
		return 0, fmt.Errorf("Interpolation ratio must be between 0 and 1, but got: %f", t)
	}
	f := t*float64(x) + (1-t)*float64(y)
	return uint8(f), nil
}
