package main

import (
	"fmt"
	"os"
)

func Echo(args []string) string {
	for idx, arg := range args {
		fmt.Printf("%d: %s\n", idx, arg)
	}
	return ""
}

func main() {
	Echo(os.Args)
}
