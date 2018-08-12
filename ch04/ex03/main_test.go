package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	seeds := [][5]int{
		[5]int{0, 1, 2, 3, 4},
		[5]int{100, 102, 102, 103, 101},
	}

	expected := [][5]int{
		[5]int{4, 3, 2, 1, 0},
		[5]int{101, 103, 102, 102, 100},
	}

	for idx, seed := range seeds {
		reverse(&seed)
		if seed != expected[idx] {
			t.Errorf("unexpected result. expected: %v, but got: %v", expected[idx], seed)
		}
	}
}
