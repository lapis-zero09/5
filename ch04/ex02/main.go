package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
)

func main() {
	f := flag.Int("type", 256, "sha size. support [256, 384, 512]")
	flag.Parse()

	var h hash.Hash
	switch *f {
	case 256:
		h = sha256.New()
	case 384:
		h = sha512.New384()
	case 512:
		h = sha512.New()
	default:
		fmt.Printf("unsupported type: %v", *f)
	}

	var s string
	fmt.Scan(&s)
	fmt.Printf("input: %s\n", s)
	h.Write([]byte(s))
	fmt.Printf("hex: %s\n", hex.EncodeToString(h.Sum(nil)))
}
