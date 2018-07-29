package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 100, 100, 100}
		}
	}
	return color.Black
}

func meanColor(cs ...color.Color) color.Color {
	var meanR, meanG, meanB, meanA uint32
	for _, c := range cs {
		r, g, b, a := c.RGBA()
		meanR += r
		meanG += g
		meanB += b
		meanA += a
	}
	s := uint32(len(cs))
	meanR /= s
	meanG /= s
	meanB /= s
	meanA /= s
	return color.RGBA{uint8(meanR), uint8(meanG), uint8(meanB), uint8(meanA)}
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	dy, dx := (0.5/height)*(ymax-ymin), (0.5/width)*(xmax-xmin)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z1 := complex(x, y)
			z2 := complex(x+dx, y)
			z3 := complex(x, y+dy)
			z4 := complex(x+dx, y+dy)

			img.Set(px, py, meanColor(mandelbrot(z1), mandelbrot(z2), mandelbrot(z3), mandelbrot(z4)))
		}
	}
	png.Encode(os.Stdout, img)
}
