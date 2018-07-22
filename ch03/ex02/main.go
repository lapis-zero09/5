package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	// r := math.Hypot(x, y)
	// return math.Sin(r) / r
	// return (x*x - x*y) / 400
	return (x*x - y*y)
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func checkInfAndPrint(xys []float64) {
	for _, val := range xys {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			return
		}
	}
	fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", xys[0], xys[1], xys[2], xys[3], xys[4], xys[5], xys[6], xys[7])
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			xys := make([]float64, 8)
			xys[0], xys[1] = corner(i+1, j)
			xys[2], xys[3] = corner(i, j)
			xys[4], xys[5] = corner(i, j+1)
			xys[6], xys[7] = corner(i+1, j+1)
			checkInfAndPrint(xys)
		}
	}
	fmt.Println("</svg>")
}
