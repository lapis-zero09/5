package charcount

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCharCount(t *testing.T) {
	ts := []struct {
		input   []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			input:   nil,
			counts:  map[rune]int{},
			utflen:  []int{0, 0, 0, 0, 0},
			invalid: 0,
		},
		{
			input:   []byte("a"),
			counts:  map[rune]int{'a': 1},
			utflen:  []int{0, 1, 0, 0, 0},
			invalid: 0,
		},
		{
			input:   []byte("é"),
			counts:  map[rune]int{'é': 1},
			utflen:  []int{0, 0, 1, 0, 0},
			invalid: 0,
		},
		{
			input:   []byte("あ"),
			counts:  map[rune]int{'あ': 1},
			utflen:  []int{0, 0, 0, 1, 0},
			invalid: 0,
		},
		{
			input:   []byte("𩸽"),
			counts:  map[rune]int{'𩸽': 1},
			utflen:  []int{0, 0, 0, 0, 1},
			invalid: 0,
		},
		{
			input:   []byte("あa"),
			counts:  map[rune]int{'あ': 1, 'a': 1},
			utflen:  []int{0, 1, 0, 1, 0},
			invalid: 0,
		},
		{
			input:   []byte{0x80, 0x80},
			counts:  map[rune]int{},
			utflen:  []int{0, 0, 0, 0, 0},
			invalid: 2,
		},
		{
			input:   []byte("^ - ^"),
			counts:  map[rune]int{32: 2, 45: 1, 94: 2},
			utflen:  []int{0, 5, 0, 0, 0},
			invalid: 0,
		},
		{
			input:   []byte("اللغة العربية"),
			counts:  map[rune]int{1604: 3, 1577: 2, 1610: 1, 1576: 1, 1575: 2, 1594: 1, 32: 1, 1593: 1, 1585: 1},
			utflen:  []int{0, 1, 12, 0, 0},
			invalid: 0,
		},
		{
			input:   []byte("😄🇯🇵"),
			counts:  map[int32]int{128516: 1, 127471: 1, 127477: 1},
			utflen:  []int{0, 0, 0, 0, 3},
			invalid: 0,
		},
	}

	for _, tc := range ts {
		counts, utflen, invalid, err := CharCount(bytes.NewReader(tc.input))
		if err != nil {
			t.Error(err)
			continue
		}
		if got, expected := counts, tc.counts; !reflect.DeepEqual(got, expected) {
			t.Errorf("unexpected counts. expected: %#v, but got: %#v", expected, got)
		}
		if got, expected := utflen, tc.utflen; !reflect.DeepEqual(got, expected) {
			t.Errorf("unexpected utflen. expected: %#v, but got: %#v", expected, got)
		}
		if got, expected := invalid, tc.invalid; got != expected {
			t.Errorf("unexpected invalid. expected: %#v, but got: %#v", expected, got)
		}
	}
}
