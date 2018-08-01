package newton

import (
	"image/color"
	"math/cmplx"
)

const (
	iterations = 37
	contrast   = 7
)

func NewtonComplex64(x, y float64) color.Color {
	z := complex64(complex(x, y))
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(complex128(z*z*z*z-1)) < 1e-6 {
			return color.RGBA{255 - contrast*i, 100, 100, 100}
		}
	}
	return color.Black
}

func NewtonComplex128(x, y float64) color.Color {
	z := complex128(complex(x, y))
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{255 - contrast*i, 100, 100, 100}
		}
	}
	return color.Black
}
