package newton

import (
	"image/color"
	"math/big"
)

type bigRatComplex struct {
	real, imag *big.Rat
}

func bigRatMul(bfc1, bfc2 *bigRatComplex) *bigRatComplex {
	return &bigRatComplex{
		new(big.Rat).Sub(new(big.Rat).Mul(bfc1.real, bfc2.real), new(big.Rat).Mul(bfc1.imag, bfc2.imag)),
		new(big.Rat).Add(new(big.Rat).Mul(bfc1.real, bfc2.imag), new(big.Rat).Mul(bfc1.imag, bfc2.real)),
	}
}

func bigRatSub(bfc1, bfc2 *bigRatComplex) *bigRatComplex {
	return &bigRatComplex{
		new(big.Rat).Sub(bfc1.real, bfc2.real),
		new(big.Rat).Sub(bfc1.imag, bfc2.imag),
	}
}

func bigRatDiv(bfc1, bfc2 *bigRatComplex) *bigRatComplex {
	conjugateComplex := &bigRatComplex{
		bfc2.real,
		new(big.Rat).Mul(big.NewRat(-1, 1), bfc2.imag),
	}
	numerator := bigRatMul(bfc1, conjugateComplex)
	denominator := bigRatMul(bfc2, conjugateComplex)
	return &bigRatComplex{
		new(big.Rat).Quo(numerator.real, new(big.Rat).Add(denominator.real, denominator.imag)),
		new(big.Rat).Quo(numerator.imag, new(big.Rat).Add(denominator.real, denominator.imag)),
	}
}

func bigRatDistance(bfc *bigRatComplex) *big.Rat {
	return new(big.Rat).Add(new(big.Rat).Mul(bfc.real, bfc.real), new(big.Rat).Mul(bfc.imag, bfc.imag))
}

func bigRatCmp(distance *big.Rat, eps float64) bool {
	if distance.Cmp(new(big.Rat).SetFloat64(eps*eps)) == -1 {
		return true
	}
	return false
}

func NewtonBigRat(x, y float64) color.Color {
	if x == 0 && y == 0 {
		return color.Black
	}
	z := &bigRatComplex{
		new(big.Rat).SetFloat64(x),
		new(big.Rat).SetFloat64(y),
	}
	one := &bigRatComplex{
		big.NewRat(1, 1),
		big.NewRat(0, 1),
	}
	four := &bigRatComplex{
		big.NewRat(4, 1),
		big.NewRat(0, 1),
	}

	for i := uint8(0); i < iterations; i++ {
		// z -= (z - 1/(z*z*z)) / 4
		z = bigRatSub(z, bigRatDiv(bigRatSub(z, bigRatDiv(one, bigRatMul(z, bigRatMul(z, z)))), four))
		cond := bigRatSub(bigRatMul(z, bigRatMul(z, bigRatMul(z, z))), one)
		if bigRatCmp(bigRatDistance(cond), 1e-6) {
			return color.RGBA{255 - contrast*i, 100, 100, 100}
		}
	}
	return color.Black
}
