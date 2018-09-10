package main

import (
	"bytes"
	"fmt"
	"io"
)

type countingWriter struct {
	w     io.Writer
	nbyte int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.nbyte += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w: w}
	return cw, &cw.nbyte
}

func main() {
	var b bytes.Buffer
	cw, _ := CountingWriter(&b)
	fmt.Println(cw)

	cw.Write([]byte("test"))
	fmt.Println(cw)

	cw.Write([]byte("aaa"))
	fmt.Println(cw)

}
