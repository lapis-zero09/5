package main

import (
	"testing"
)

func equals(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRemoveSameAdjacent(t *testing.T) {
	seeds := [][]int{
		[]int{0, 1, 2, 3, 4},
		[]int{0, 0, 0, 1, 0},
		[]int{1, 1, 1, 1, 1},
	}

	expected := [][]int{
		[]int{0, 1, 2, 3, 4},
		[]int{0, 1, 0},
		[]int{1},
	}

	for idx, seed := range seeds {
		ans := removeSameAdjacent(seed)
		if !equals(expected[idx], ans) {
			t.Errorf("unexpected result. expected: %v, but got: %v", expected[idx], seed)
		}
	}
}
