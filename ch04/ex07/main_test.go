package main

import "testing"

func TestReverseUTF8(t *testing.T) {
	seeds := []string{
		"abcde",
		"アイウエオ",
		"🤖 💀 👵 👩 👧 👶",
	}

	expected := []string{
		"edcba",
		"オエウイア",
		"👶 👧 👩 👵 💀 🤖",
	}

	for idx, ex := range expected {
		b := []byte(seeds[idx])
		reverseUTF8(b)
		if string(b) != ex {
			t.Errorf("unexpected result. expected: %v, but got: %v", ex, string(b))
		}
	}
}
