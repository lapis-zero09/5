package main

import (
	"io/ioutil"
	"testing"

	"github.com/5/ch03/ex08/newton"
)

func BenchmarkNewtonComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plot(newton.NewtonComplex64, ioutil.Discard)
	}
}

func BenchmarkNewtonComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plot(newton.NewtonComplex128, ioutil.Discard)
	}
}

func BenchmarkNewtonBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plot(newton.NewtonBigFloat, ioutil.Discard)
	}
}

// func BenchmarkNewtonBigRat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		plot(newton.NewtonBigRat, ioutil.Discard)
// 	}
// }
