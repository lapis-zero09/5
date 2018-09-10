package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	ts := []struct {
		text     string
		expected [2]int
	}{
		{
			text:     "",
			expected: [2]int{0, 0},
		},
		{
			text:     "\n",
			expected: [2]int{1, 1},
		},
		{
			text:     "\n\n\n",
			expected: [2]int{1, 3},
		},
		{
			text:     "\n\n\ntest",
			expected: [2]int{1, 4},
		},
		{
			text:     "test",
			expected: [2]int{1, 1},
		},
		{
			text:     "foo\tbar\nfoo\nbaz\nbar",
			expected: [2]int{5, 4},
		},
		{
			text:     "this is test sentence.\tbar\nfoo\nbaz\nbar",
			expected: [2]int{8, 4},
		},
	}

	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {
			var c Counter
			c.Write([]byte(tc.text))

			if c.wordsCount != tc.expected[0] {
				t.Errorf("unexpected wordsCount result. expected: %d, but got: %d", tc.expected[0], c.wordsCount)
			}
			if c.linesCount != tc.expected[1] {
				t.Errorf("unexpected linesCount result. expected: %d, but got: %d", tc.expected[1], c.linesCount)
			}

		})
	}
}
