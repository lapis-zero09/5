package main

import "testing"

func benchmark(b *testing.B, size int, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			f(0x1234567890ABCDEF)
		}
	}
}

// table
func BenchmarkPopCount1(b *testing.B) {
	benchmark(b, 1, PopCount)
}

func BenchmarkPopCount10(b *testing.B) {
	benchmark(b, 10, PopCount)
}

func BenchmarkPopCount100(b *testing.B) {
	benchmark(b, 100, PopCount)
}

func BenchmarkPopCount1000(b *testing.B) {
	benchmark(b, 1000, PopCount)
}

func BenchmarkPopCount10000(b *testing.B) {
	benchmark(b, 10000, PopCount)
}

// loop
func BenchmarkPopCountLoop1(b *testing.B) {
	benchmark(b, 1, PopCountLoop)
}

func BenchmarkPopCountLoop10(b *testing.B) {
	benchmark(b, 10, PopCountLoop)
}

func BenchmarkPopCountLoop100(b *testing.B) {
	benchmark(b, 100, PopCountLoop)
}

func BenchmarkPopCountLoop1000(b *testing.B) {
	benchmark(b, 1000, PopCountLoop)
}

func BenchmarkPopCountLoop10000(b *testing.B) {
	benchmark(b, 10000, PopCountLoop)
}

// bitshift
func BenchmarkPopCountBitShift1(b *testing.B) {
	benchmark(b, 1, PopCountBitShift)
}

func BenchmarkPopCountBitShift10(b *testing.B) {
	benchmark(b, 10, PopCountBitShift)
}

func BenchmarkPopCountBitShift100(b *testing.B) {
	benchmark(b, 100, PopCountBitShift)
}

func BenchmarkPopCountBitShift1000(b *testing.B) {
	benchmark(b, 1000, PopCountBitShift)
}

func BenchmarkPopCountBitShift10000(b *testing.B) {
	benchmark(b, 10000, PopCountBitShift)
}

// clear
func BenchmarkPopCountClear1(b *testing.B) {
	benchmark(b, 1, PopCountClear)
}

func BenchmarkPopCountClear10(b *testing.B) {
	benchmark(b, 10, PopCountClear)
}

func BenchmarkPopCountClear100(b *testing.B) {
	benchmark(b, 100, PopCountClear)
}

func BenchmarkPopCountClear1000(b *testing.B) {
	benchmark(b, 1000, PopCountClear)
}
func BenchmarkPopCountClear10000(b *testing.B) {
	benchmark(b, 10000, PopCountClear)
}
