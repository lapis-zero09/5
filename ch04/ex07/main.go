package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func firstRune2Back(b []byte, firstRuneSize int) {
	reverse(b[:firstRuneSize])
	reverse(b[firstRuneSize:])
	reverse(b)
}

func reverseUTF8(b []byte) {
	if len(b) <= 1 {
		return
	}
	_, firstRuneSize := utf8.DecodeRune(b)
	firstRune2Back(b, firstRuneSize)
	reverseUTF8(b[:len(b)-firstRuneSize])
}

func main() {
	s := []byte("🤖 💀 👵 👩 👧 👶")

	fmt.Printf("original: %s\n", string(s))
	reverseUTF8(s)
	fmt.Printf("reversed: %s\n", string(s))
}
