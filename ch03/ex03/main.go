package main

import (
	"fmt"
	"math"
)

const (
	width, height = 700, 420            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	minZ          = -.6
	maxZ          = .45
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	return (x*x - 2*y*y) / 500
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func checkInfAndPrint(xys []float64, r, b int) {
	for _, val := range xys {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			return
		}
	}

	fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgba(%d,0,%d,0.8)' stloke='rgba(%d,0,%d,0.8)'/>\n", xys[0], xys[1], xys[2], xys[3], xys[4], xys[5], xys[6], xys[7], r, b, r, b)
}

func main() {
	var az, bz, cz, dz float64
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {

		for j := 0; j < cells; j++ {
			xys := make([]float64, 8)
			xys[0], xys[1], az = corner(i+1, j)
			xys[2], xys[3], bz = corner(i, j)
			xys[4], xys[5], cz = corner(i, j+1)
			xys[6], xys[7], dz = corner(i+1, j+1)
			z := (az + bz + cz + dz) / 4
			r := int(255 * (z - minZ) / (maxZ - minZ))
			b := 255 - r
			checkInfAndPrint(xys, r, b)
		}
	}
	fmt.Println("</svg>")
}
