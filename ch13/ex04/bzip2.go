package bzip

import (
	"io"
	"log"
	"os/exec"
)

type writer struct {
	cmd *exec.Cmd
	wc  io.WriteCloser
}

func NewWriter(out io.Writer) io.WriteCloser {
	cmd := exec.Command("/usr/bin/bzip2")
	cmd.Stdout = out
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()
	return &writer{
		cmd: cmd,
		wc:  stdin,
	}
}

func (w *writer) Write(data []byte) (int, error) {
	return w.wc.Write(data)
}

func (w *writer) Close() error {
	w.wc.Close()
	return w.cmd.Wait()
}
