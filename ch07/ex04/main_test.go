package main

import (
	"bytes"
	"io"
	"testing"
)

func TestNewReader(t *testing.T) {
	ts := []string{
		"",
		"\n",
		"test",
		"\n\n\ntest",
		"アイウエオ",
	}

	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {
			r := NewReader(tc)
			var b bytes.Buffer
			n, err := io.Copy(&b, r)
			if err != nil {
				t.Error(err)
			}

			if int(n) != len(tc) {
				t.Errorf("unexpected size. expected: %d, but got: %d", len(tc), int(n))
			}
			if b.String() != tc {
				t.Errorf("unexpected result. expected: %s, but got: %s", tc, b.String())
			}

		})
	}
}
