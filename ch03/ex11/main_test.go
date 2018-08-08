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
		"+2443534",
		"132435466.7654",
		"-5644.545332454",
		"-876543245675432354657643",
	}

	expected := []string{
		"123",
		"123,456",
		"12,345,678",
		"",
		"1,243,54i,yuy,tre,gbf,dsb,gfd",
		"g,trv,faf",
		"+2,443,534",
		"132,435,466.7654",
		"-5,644.545332454",
		"-876,543,245,675,432,354,657,643",
	}

	for idx, s := range seeds {
		if g := comma(s); g != expected[idx] {
			t.Errorf("unexpected result. expected: %v, but got: %v", expected[idx], g)
		}
	}
}
