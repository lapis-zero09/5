package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteString(s[:1])
		s = s[1:]
	}

	flt := ""
	if idx := strings.IndexByte(s, '.'); idx >= 0 {
		s, flt = s[:idx], s[idx:]

	}

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
	if len(flt) > 0 {
		buf.WriteString(flt)
	}
	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}
