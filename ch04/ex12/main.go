package main

import (
	"fmt"
	"os"

	"github.com/lapis-zero09/5/ch04/ex12/xkcd"
)

func printXKCD(xkcd *xkcd.XKCD) {
	fmt.Println("---------------------------------\n")
	fmt.Printf("#%-5d %s\n%s\n",
		xkcd.Num, xkcd.Title, xkcd.Transcript)
	fmt.Println("---------------------------------\n")
}

func printXKCDs(xkcds []*xkcd.XKCD) {
	for _, xkcd := range xkcds {
		printXKCD(xkcd)
	}
}

func main() {
	// get xkcd from https://xkcd.com
	// xkcd.GetAllXKCDs()

	// check all xkcds in db
	// xkcds := xkcd.GetXKCDs()
	// printXKCDs(xkcds)

	// use query
	xkcds, err := xkcd.SearchXKCD(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("hits: %d\n", len(xkcds))
	printXKCDs(xkcds)
}
