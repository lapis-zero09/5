package main

import (
	"fmt"
	"os"
	"strings"
)

func Echo(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	ret := Echo(os.Args)
	fmt.Println(ret)
}
