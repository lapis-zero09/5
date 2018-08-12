package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countBitDiff(x, y []byte) int {
	diff := 0
	for idx := range x {
		diff += int(pc[x[idx]^y[idx]])
	}
	return diff
}

func main() {
	a := sha256.New().Sum([]byte("aaaaaaaaaa"))
	b := sha256.New().Sum([]byte("aaaaaaaabb"))
	fmt.Printf("a:%x\nb:%x\n", a, b)
	fmt.Printf("countBitDiff: %d\n", countBitDiff(a, b))
}
