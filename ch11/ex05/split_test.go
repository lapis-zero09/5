package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	ts := []struct {
		s, sep   string
		expected int
	}{
		{"a:b:c", ":", 3},
		{"afds,a re,g re,ger,gre f,f", ",", 6},
		{"", "\n", 1},
		{"1 4 2 4 4 2 32f f  sav ", " ", 11},
	}

	for _, tc := range ts {
		words := strings.Split(tc.s, tc.sep)
		if got, expected := len(words), tc.expected; got != expected {
			t.Errorf("Split(%q, %q) returned %d words, expected %d", tc.s, tc.sep, got, expected)
		}
	}
}
