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

func TestRotate(t *testing.T) {
	seed := []int{0, 1, 2, 3, 4}
	expected := []int{4, 0, 1, 2, 3}
	rotate(seed, 1)
	if !equals(seed, expected) {
		t.Errorf("unexpected result. expected: %v, but got: %v", expected, seed)
	}

	seed = []int{0, 0, 0, 1, 0}
	expected = []int{0, 1, 0, 0, 0}
	rotate(seed, 3)
	if !equals(seed, expected) {
		t.Errorf("unexpected result. expected: %v, but got: %v", expected, seed)
	}
}
