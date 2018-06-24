package main

import (
	"fmt"
	"os"
	"strings"
)

func efficientEcho(args []string) string {
	return strings.Join(args[1:], " ")
}

func inefficientEcho(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func main() {
	fmt.Println("efficientEcho")
	ret := efficientEcho(os.Args)
	fmt.Println(ret)

	fmt.Println("inefficientEcho")
	ret = inefficientEcho(os.Args)
	fmt.Println(ret)
}
