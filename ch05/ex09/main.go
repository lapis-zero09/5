package main

import (
	"fmt"
	"regexp"
)

var r = regexp.MustCompile(`\$\w*`)

func expand(s string, f func(string) string) string {
	if f == nil {
		return s
	}
	return r.ReplaceAllStringFunc(s, func(s string) string {
		return f(s[1:])
	})
}

func main() {
	s := "test $test brneagbr $ aaa $explor"
	fmt.Printf("before:\t%s\n", s)
	s = expand(s, func(s string) string {
		return s + "ed"
	})
	fmt.Printf("after :\t%s\n", s)
}
