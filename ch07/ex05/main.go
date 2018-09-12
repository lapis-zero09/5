package main

import (
	"io"
)

type Reader struct {
	r io.Reader
	n int64 // current reading index
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &Reader{
		r: r,
		n: n,
	}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.n {
		p = p[0:r.n]
	}

	n, err = r.r.Read(p)
	r.n -= int64(n)
	return
}
