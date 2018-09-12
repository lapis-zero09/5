package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	ts := []struct {
		text     string
		limit    int64
		expected string
	}{
		{
			text:     "",
			limit:    0,
			expected: "",
		},
		{
			text:     "",
			limit:    10,
			expected: "",
		},
		{
			text:     "\n\n\ntest",
			limit:    int64(len([]byte("\n\n\nt"))),
			expected: "\n\n\nt",
		},
		{
			text:     "this is test sentence.\tbar\nfoo\nbaz\nbar",
			limit:    int64(len([]byte("this is test sentence.\tbar\nfoo\nbaz\nbar"))),
			expected: "this is test sentence.\tbar\nfoo\nbaz\nbar",
		},
	}

	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {
			r := strings.NewReader(tc.text)
			lr := LimitReader(r, tc.limit)

			var b bytes.Buffer
			if _, err := io.Copy(&b, lr); err != nil {
				t.Error(err)
			}

			if b.String() != tc.expected {
				t.Errorf("unexpected result. expected: %s, but got: %s", tc.expected, b.String())
			}

		})
	}
}
