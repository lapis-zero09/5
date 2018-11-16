package main_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/lapis-zero09/5/ch11/ex07/intset"
	"github.com/lapis-zero09/5/ch11/ex07/mapintset"
)

var seed = time.Now().UTC().UnixNano()

func benchMapIntSetAdd(b *testing.B, size int) {
	rand.Seed(seed)
	for i := 0; i < b.N; i++ {
		mapIntSet := &mapintset.MapIntSet{}
		for j := 0; j < size; j++ {
			mapIntSet.Add(rand.Intn(1000000))
		}
	}
}

func BenchmarkMapIntSetAdd1(b *testing.B) {
	benchMapIntSetAdd(b, 1)
}

func BenchmarkMapIntSetAdd10(b *testing.B) {
	benchMapIntSetAdd(b, 10)
}

func BenchmarkMapIntSetAdd100(b *testing.B) {
	benchMapIntSetAdd(b, 100)
}

func BenchmarkMapIntSetAdd1000(b *testing.B) {
	benchMapIntSetAdd(b, 1000)
}

func BenchmarkMapIntSetAdd10000(b *testing.B) {
	benchMapIntSetAdd(b, 10000)
}

func BenchmarkMapIntSetAdd100000(b *testing.B) {
	benchMapIntSetAdd(b, 100000)
}

func BenchmarkMapIntSetAdd1000000(b *testing.B) {
	benchMapIntSetAdd(b, 1000000)
}

func benchIntSetAdd(b *testing.B, size int) {
	rand.Seed(seed)
	for i := 0; i < b.N; i++ {
		intSet := &intset.IntSet{}
		for j := 0; j < size; j++ {
			intSet.Add(rand.Intn(1000000))
		}
	}
}

func BenchmarkIntSetAdd1(b *testing.B) {
	benchIntSetAdd(b, 1)
}

func BenchmarkIntSetAdd10(b *testing.B) {
	benchIntSetAdd(b, 10)
}

func BenchmarkIntSetAdd100(b *testing.B) {
	benchIntSetAdd(b, 100)
}

func BenchmarkIntSetAdd1000(b *testing.B) {
	benchIntSetAdd(b, 1000)
}

func BenchmarkIntSetAdd10000(b *testing.B) {
	benchIntSetAdd(b, 10000)
}

func BenchmarkIntSetAdd100000(b *testing.B) {
	benchIntSetAdd(b, 100000)
}

func BenchmarkIntSetAdd1000000(b *testing.B) {
	benchIntSetAdd(b, 1000000)
}
