package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	tcpaddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdout)
	conn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
