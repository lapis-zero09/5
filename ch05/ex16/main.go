package main

import (
	"fmt"
	"strings"
)

func Join(sep string, a ...string) string {
	return strings.Join(a, sep)
}

func main() {
	values := []string{"a", "b", "c"}
	fmt.Println(Join(",", values...))
}
