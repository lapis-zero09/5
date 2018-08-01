package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"

	"github.com/5/ch03/ex08/newton"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func plot(method func(x, y float64) color.Color, w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, method(x, y))
		}
	}
	png.Encode(w, img)
}

func main() {
	plot(newton.NewtonComplex64, os.Stdout)
	// plot(newton.NewtonComplex128, os.Stdout)
	// plot(newton.NewtonBigFloat, os.Stdout)
	// plot(newton.NewtonBigRat, os.Stdout)
}
