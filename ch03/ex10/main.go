package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	var buf bytes.Buffer
	for len(s) > 3 {
		size := len(s) % 3
		if size == 0 {
			size = 3
		}
		buf.WriteString(s[:size])
		buf.WriteString(",")
		s = s[size:]
	}
	buf.WriteString(s)
	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}
