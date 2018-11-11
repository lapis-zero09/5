package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"github.com/lapis-zero09/5/ch10/ex02/archive"
)

func init() {
	archive.RegFmt("zip", 0, []byte{'P', 'K', 0x05, 0x06}, Unarchive)
}

type Archive struct {
	zip *zip.Reader
	cur int
}

func (a *Archive) Next() (*archive.File, error) {
	if a.cur >= len(a.zip.File) {
		return nil, io.EOF
	}

	f := a.zip.File[a.cur]
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}

	return &archive.File{
		Name: f.Name,
		Body: rc,
	}, nil
}

type sizeReaderAt interface {
	io.ReaderAt
	Size() int64
}

type File struct {
	*os.File
	size int64
}

func (f *File) Size() int64 { return f.size }

func asSizeReaderAt(r io.Reader) (sizeReaderAt, error) {
	if f, ok := r.(*os.File); ok {
		stat, err := f.Stat()
		if err != nil {
			return nil, err
		}
		return &File{
			File: f,
			size: stat.Size(),
		}, nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

func Unarchive(r io.Reader) (archive.Archive, error) {
	sra, err := asSizeReaderAt(r)
	if err != nil {
		return nil, err
	}

	zr, err := zip.NewReader(sra, sra.Size())
	if err != nil {
		return nil, err
	}

	return &Archive{zip: zr}, nil
}
