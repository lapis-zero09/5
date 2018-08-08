package main

import "testing"

func TestComma(t *testing.T) {
	seeds := []string{
		"123",
		"123456",
		"12345678",
		"",
		"124354iyuytregbfdsbgfd",
		"gtrvfaf",
	}

	expected := []string{
		"123",
		"123,456",
		"12,345,678",
		"",
		"1,243,54i,yuy,tre,gbf,dsb,gfd",
		"g,trv,faf",
	}

	for idx, s := range seeds {
		if g := comma(s); g != expected[idx] {
			t.Errorf("unexpected result. expected: %v, but got: %v", expected[idx], g)
		}
	}
}
