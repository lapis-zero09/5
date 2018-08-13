package main

import (
	"fmt"
)

func spaceRange(b []byte) int {
	if len(b) <= 1 {
		return 0
	}
	// ref: http://lwp.interglacial.com/appf_01.htm
	switch b[0] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xa0:
		return 1
	case 0xe2:
		if len(b) <= 3 {
			return 0
		}

		if b[1] == 0x80 && (b[2] == 0x82 || b[2] == 0x83 || b[2] == 0x89) {
			return 3
		}
		return 0

	case 0xe3:
		if len(b) <= 3 {
			return 0
		}
		if b[1] == 0x80 && b[2] == 0x80 {
			return 3
		}
		return 0
	default:
		return 0
	}
}

func trimBlank(b []byte) []byte {
	trimmedByte := b[:0]
	skipFlag := false
	for idx := 0; idx < len(b); idx++ {
		r := spaceRange(b[idx:])
		if r > 0 {
			if !skipFlag {
				trimmedByte = append(trimmedByte, ' ')
				skipFlag = true
			}
			idx += r - 1
			continue
		}
		trimmedByte = append(trimmedByte, b[idx])
		skipFlag = false
	}
	return trimmedByte
}

func main() {
	s := "a　 b b s f f e  gref \n asfd \t fds\n"
	fmt.Printf("org: %s", s)
	fmt.Printf("trimmed: %s", string(trimBlank([]byte(s))))

	s = "asfd \t fds\t"
	fmt.Printf("org: %s", s)
	fmt.Printf("trimmed: %s", string(trimBlank([]byte(s))))

	s = "asfd \t fds\v"
	fmt.Printf("org: %s", s)
	fmt.Printf("trimmed: %s", string(trimBlank([]byte(s))))
}
