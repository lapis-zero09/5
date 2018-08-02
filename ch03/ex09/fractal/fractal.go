package fractal

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{255 - contrast*i, 100, 100, 100}
		}
	}
	return color.Black
}

func Fractal(out io.Writer, params map[string]int) {
	x := float64(params["x"])
	y := float64(params["y"])
	xmax, ymax := x, y
	xmin, ymin := -1.*x, -1.*y
	width, height := params["magnification"], params["magnification"]

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(out, img)
}
