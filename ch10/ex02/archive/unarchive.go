package archive

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

func Unarchive(r io.Reader) (Archive, string, error) {
	r, p := asPeeker(r)
	f := sniff(p)
	if f.unarchive == nil {
		return nil, "", errors.New("unknown format")
	}

	a, err := f.unarchive(r)
	if err != nil {
		return nil, f.name, err
	}
	return a, f.name, nil
}

type Archive interface {
	Next() (*File, error)
}

type File struct {
	Name string
	Body io.ReadCloser
}

type format struct {
	name      string
	offset    int
	magic     []byte
	unarchive func(io.Reader) (Archive, error)
}

var formats []format

func RegFmt(name string, offset int, magic []byte, unarchive func(io.Reader) (Archive, error)) {
	formats = append(formats, format{
		name:      name,
		offset:    offset,
		magic:     magic,
		unarchive: unarchive,
	})
}

type peeker interface {
	Peek(int) ([]byte, error)
}

func (p *atPeeker) Peek(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := p.ReadAt(b, 0)
	return b, err
}

type atPeeker struct {
	io.ReaderAt
}

func asPeeker(r io.Reader) (io.Reader, peeker) {
	switch t := r.(type) {
	case peeker:
		return r, t
	case io.ReaderAt:
		return r, &atPeeker{t}
	default:
		b := bufio.NewReader(r)
		return b, b
	}
}

func sniff(p peeker) format {
	for _, f := range formats {
		b, err := p.Peek(f.offset + len(f.magic))
		if err == nil && bytes.Equal(b[f.offset:], f.magic) {
			return f
		}
	}
	return format{}
}
