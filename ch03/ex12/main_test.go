package main

import (
	"testing"
)

func TestString2byteCount(t *testing.T) {
	seeds := []string{
		"",
		"hello world",
		"this is test sentence01.",
	}

	expectedByteCnts := []map[byte]int{
		{},
		{
			100: 1,
			104: 1,
			101: 1,
			108: 3,
			111: 2,
			32:  1,
			119: 1,
			114: 1,
		},
		{
			116: 4,
			104: 1,
			105: 2,
			115: 4,
			32:  3,
			101: 4,
			110: 2,
			46:  1,
			99:  1,
			48:  1,
			49:  1,
		},
	}

	for idx, expected := range expectedByteCnts {
		cnt := string2byteCount(seeds[idx])

		// expected -> string2byteCount
		for key, val := range expected {
			v, ok := cnt[key]
			if !ok || val != v {
				t.Errorf("unexpected result. expected: %v, but got: %v", expected, cnt)
			}
		}

		// string2byteCount -> expected
		for key, val := range cnt {
			v, ok := expected[key]
			if !ok || val != v {
				t.Errorf("unexpected result. expected: %v, but got: %v", expected, cnt)
			}
		}
	}
}

func TestIsAnagram(t *testing.T) {
	seeds := [][]string{
		[]string{"", ""},
		[]string{"abcd", "abcd"},
		[]string{"abcd", "abc"},
		[]string{"abd", "dab"},
		[]string{"asdf", "fdsa"},
		[]string{"qwer", "abcd"},
		[]string{"あいうえお亜", "うお絵いう"},
		[]string{"んあ", "あん"},
	}

	expected := []bool{
		true,
		true,
		false,
		true,
		true,
		false,
		false,
		true,
	}

	for idx, seed := range seeds {
		ans := isAnagram(seed[0], seed[1])
		if ex := expected[idx]; ans != ex {
			t.Errorf("unexpected result. expected: %v, but got: %v", ex, ans)
		}
	}
}
