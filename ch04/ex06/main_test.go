package main

import (
	"testing"
)

func TestSpaceRange(t *testing.T) {
	seeds := [][]byte{
		[]byte{},
		[]byte{0xe2},
		[]byte{'\t', 0x00},
		[]byte{0xe2, 0x80},
		[]byte{0xe2, 0x80, 0x81, 0x00},
		[]byte{0xe2, 0x80, 0x82, 0x00},
		[]byte{0xe2, 0x80, 0x83, 0x00},
		[]byte{0xe2, 0x80, 0x89, 0x00},
		[]byte{0xe3, 0x80},
		[]byte{0xe3, 0x80, 0x80, 0x00},
		[]byte{0xe3, 0x80, 0x81, 0x00},
	}

	expected := []int{
		0,
		0,
		1,
		0,
		0,
		3,
		3,
		3,
		0,
		3,
		0,
	}

	for idx, seed := range seeds {
		if ans := spaceRange(seed); expected[idx] != ans {
			t.Errorf("unexpected result. expected: %v, but got: %v", expected[idx], ans)
		}
	}
}

func TestTrimBlank(t *testing.T) {
	seeds := []string{
		"a b",
		"a  b",
		"a    b",
		"a 　  b",
		"a 　 \t b\t",
		"a 　 \v b\v",
		"a 　 \f b\f",
		"a 　 \r b\r",
		"a 　 \n b\n",
	}

	expected := []string{
		"a b",
		"a b",
		"a b",
		"a b",
		"a b\t",
		"a b\v",
		"a b\f",
		"a b\r",
		"a b\n",
	}

	for idx, seed := range seeds {
		if ans := string(trimBlank([]byte(seed))); expected[idx] != ans {
			t.Errorf("unexpected result. expected: %v, but got: %v", expected[idx], ans)
		}
	}
}
