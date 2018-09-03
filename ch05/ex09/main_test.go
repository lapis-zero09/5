package main

import (
	"testing"
)

func makeHTML(body string) string {
	return "<html><body>" + body + "</body</html>"
}

func TestElementByID(t *testing.T) {
	ts := []struct {
		seed     string
		f        func(string) string
		expected string
	}{
		{
			seed:     "",
			f:        func(s string) string { return s + "1" },
			expected: "",
		},
		{
			seed:     "",
			f:        nil,
			expected: "",
		},
		{
			seed:     "test $test aaaa $aaaa",
			f:        nil,
			expected: "test $test aaaa $aaaa",
		},
		{
			seed:     "test $test aaaa $aaaa",
			f:        func(s string) string { return "(" + s + ")" },
			expected: "test (test) aaaa (aaaa)",
		},
		{
			seed:     "test $test aaaa $aaaa",
			f:        func(s string) string { return "**censored**" },
			expected: "test **censored** aaaa **censored**",
		},
	}

	for _, tc := range ts {
		t.Run("Case", func(t *testing.T) {

			s := expand(tc.seed, tc.f)
			if s != tc.expected {
				t.Errorf("unexpected result. expected: (%s), but got: (%s)", tc.expected, s)
			}
		})
	}
}
