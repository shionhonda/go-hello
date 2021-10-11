package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 200                 // number of grid cells
	xyrange       = 2.0                 // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j, sixhump)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j, sixhump)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1, sixhump)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1, sixhump)
			if err != nil {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f func(float64, float64) float64) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	// Return error if z in Inf or NaN
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0., 0., fmt.Errorf("Invalid output: %v", z)
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func sixhump(x, y float64) float64 {
	x2 := math.Pow(x, 2)
	x4 := math.Pow(x, 4)
	y2 := math.Pow(y, 2)
	return ((4-2.1*x2+x4/3)*x2 + x*y + (-4+4*y2)*y2)
}
