package tar

import (
	"archive/tar"
	"io"
	"io/ioutil"

	"github.com/lapis-zero09/5/ch10/ex02/archive"
)

func init() {
	archive.RegFmt("tar", 0x101, []byte("ustar.00"), Unarchive)
	archive.RegFmt("tar", 0x101, []byte("ustar . "), Unarchive)

}

type Archive struct {
	tar *tar.Reader
}

func (a *Archive) Next() (*archive.File, error) {
	h, err := a.tar.Next()
	if err != nil {
		return nil, err
	}

	return &archive.File{
		Name: h.Name,
		Body: ioutil.NopCloser(a.tar),
	}, nil
}

func Unarchive(r io.Reader) (archive.Archive, error) {
	t := tar.NewReader(r)
	return &Archive{tar: t}, nil
}
