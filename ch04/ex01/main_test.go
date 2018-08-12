package main

import (
	"crypto/sha256"
	"testing"
)

func TestCountBitDiff(t *testing.T) {
	seeds := [][]string{
		[]string{"", ""},
		[]string{"a", "a"},
		[]string{"X", "x"},
		[]string{"c", "C"},
		[]string{"a", "b"},
		[]string{"aaaaaaaa", "aaaaaaab"},
		[]string{"aaaaaaaa", "aaaaaabb"},
		[]string{"aaaaaaaa", "aaaaacbb"},
	}

	expected := []int{
		0,
		0,
		1,
		1,
		2,
		2,
		4,
		5,
	}

	for idx, seed := range seeds {
		a := sha256.New().Sum([]byte(seed[0]))
		b := sha256.New().Sum([]byte(seed[1]))
		if ans := countBitDiff(a, b); ans != expected[idx] {
			t.Errorf("unexpected result. expected: %v, but got: %v", expected[idx], ans)
		}
	}
}
