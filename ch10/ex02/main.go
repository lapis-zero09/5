package main

import (
	"fmt"
	"os"

	"github.com/lapis-zero09/5/ch10/ex02/archive"
	_ "github.com/lapis-zero09/5/ch10/ex02/archive/tar"
)

func main() {
	a, fname, err := archive.Unarchive(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(fname)
	fmt.Fprint(os.Stdout, a)
}
