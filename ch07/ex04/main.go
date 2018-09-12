package main

import (
	"io"
)

type Reader struct {
	s string
	i int64 // current reading index
}

func NewReader(s string) io.Reader {
	return &Reader{s: s}
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
