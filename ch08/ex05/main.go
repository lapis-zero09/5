package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

type colors struct {
	x, y int
	c    color.Color
}

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

func main() {
	parallels := 12
	pChan := make(chan *colors, parallels)
	mChan := make(chan *colors, parallels)

	for i := 0; i < parallels; i++ {
		go func() {
			for p := range pChan {
				y := float64(p.y)/height*(ymax-ymin) + ymin
				x := float64(p.x)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				mChan <- &colors{p.x, p.y, mandelbrot(z)}
			}
		}()
	}

	go func() {
		for py := 0; py < height; py++ {
			for px := 0; px < width; px++ {
				pChan <- &colors{x: px, y: py}
			}
		}
		close(pChan)
	}()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height*width; i++ {
		p := <-mChan
		img.Set(p.x, p.y, p.c)
	}

	png.Encode(os.Stdout, img)
}
