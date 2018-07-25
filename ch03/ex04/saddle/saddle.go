package saddle

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func f(x, y float64) float64 {
	return (x*x - 2*y*y) / 1000
}

func corner(i, j int, width, height, xyscale, zscale float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func checkInfAndPrint(out io.Writer, xys []float64, color string) {
	for _, val := range xys {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			return
		}
	}

	fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' stloke='%s'/>\n", xys[0], xys[1], xys[2], xys[3], xys[4], xys[5], xys[6], xys[7], color, color)
}

func Saddle(out io.Writer, params map[string]string) {
	color := params["color"]
	width, err := strconv.ParseFloat(params["width"], 64)
	if err != nil {
		width = 1000.
	}
	height, err := strconv.ParseFloat(params["height"], 64)
	if err != nil {
		height = 1000.
	}
	xyscale := width / 2 / xyrange
	zscale := height * 0.4

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%f' height='%f'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			xys := make([]float64, 8)
			xys[0], xys[1] = corner(i+1, j, width, height, xyscale, zscale)
			xys[2], xys[3] = corner(i, j, width, height, xyscale, zscale)
			xys[4], xys[5] = corner(i, j+1, width, height, xyscale, zscale)
			xys[6], xys[7] = corner(i+1, j+1, width, height, xyscale, zscale)
			checkInfAndPrint(out, xys, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}
