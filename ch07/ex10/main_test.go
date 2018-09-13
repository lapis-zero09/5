package main

import (
	"sort"
	"testing"
)

func TestNewReader(t *testing.T) {
	ts := []struct {
		s        sort.Interface
		expected bool
	}{
		{
			s:        sort.IntSlice([]int{}),
			expected: true,
		},
		{
			s:        sort.IntSlice([]int{0}),
			expected: true,
		},
		{
			s:        sort.IntSlice([]int{1, 2, 3, 2, 0}),
			expected: false,
		},
		{
			s:        sort.IntSlice([]int{1, 2, 3, 4, 3, 2, 1}),
			expected: true,
		},
	}

	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {
			if g := IsPalindrome(tc.s); g != tc.expected {
				t.Errorf("unexpected result. expected: %v, but got: %v", tc.expected, g)
			}

		})
	}
}
