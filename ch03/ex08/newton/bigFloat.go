package newton

import (
	"image/color"
	"math/big"
)

type bigFloatComplex struct {
	real, imag *big.Float
}

func bigFloatMul(bfc1, bfc2 *bigFloatComplex) *bigFloatComplex {
	return &bigFloatComplex{
		new(big.Float).Sub(new(big.Float).Mul(bfc1.real, bfc2.real), new(big.Float).Mul(bfc1.imag, bfc2.imag)),
		new(big.Float).Add(new(big.Float).Mul(bfc1.real, bfc2.imag), new(big.Float).Mul(bfc1.imag, bfc2.real)),
	}
}

func bigFloatSub(bfc1, bfc2 *bigFloatComplex) *bigFloatComplex {
	return &bigFloatComplex{
		new(big.Float).Sub(bfc1.real, bfc2.real),
		new(big.Float).Sub(bfc1.imag, bfc2.imag),
	}
}

func bigFloatDiv(bfc1, bfc2 *bigFloatComplex) *bigFloatComplex {
	conjugateComplex := &bigFloatComplex{
		bfc2.real,
		new(big.Float).Mul(big.NewFloat(-1.), bfc2.imag),
	}
	numerator := bigFloatMul(bfc1, conjugateComplex)
	denominator := bigFloatMul(bfc2, conjugateComplex)
	return &bigFloatComplex{
		new(big.Float).Quo(numerator.real, new(big.Float).Add(denominator.real, denominator.imag)),
		new(big.Float).Quo(numerator.imag, new(big.Float).Add(denominator.real, denominator.imag)),
	}
}

func bigFloatDistance(bfc *bigFloatComplex) *big.Float {
	return new(big.Float).Add(new(big.Float).Mul(bfc.real, bfc.real), new(big.Float).Mul(bfc.imag, bfc.imag))
}

func bigFloatCmp(distance *big.Float, eps float64) bool {
	if distance.Cmp(big.NewFloat(eps*eps)) == -1 {
		return true
	}
	return false
}

func NewtonBigFloat(x, y float64) color.Color {
	if x == 0 && y == 0 {
		return color.Black
	}
	z := &bigFloatComplex{
		big.NewFloat(x),
		big.NewFloat(y),
	}
	one := &bigFloatComplex{
		big.NewFloat(1),
		big.NewFloat(0),
	}
	four := &bigFloatComplex{
		big.NewFloat(4),
		big.NewFloat(0),
	}

	for i := uint8(0); i < iterations; i++ {
		// z -= (z - 1/(z*z*z)) / 4
		z = bigFloatSub(z, bigFloatDiv(bigFloatSub(z, bigFloatDiv(one, bigFloatMul(z, bigFloatMul(z, z)))), four))
		cond := bigFloatSub(bigFloatMul(z, bigFloatMul(z, bigFloatMul(z, z))), one)
		if bigFloatCmp(bigFloatDistance(cond), 1e-6) {
			return color.RGBA{255 - contrast*i, 100, 100, 100}
		}
	}
	return color.Black
}
