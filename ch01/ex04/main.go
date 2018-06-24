package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	filenameList := dup(files)
	for line, list := range filenameList {
		if len(list) > 1 {
			fmt.Printf("%s\t%v\n", line, list)
		}
	}
}

func isPresent(list []string, query string) bool {
	for _, elem := range list {
		if elem == query {
			return true
		}
	}
	return false
}

func countLines(f *os.File, filename string, filenameList map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if isPresent(filenameList[input.Text()], filename) != true {
			filenameList[input.Text()] = append(filenameList[input.Text()], filename)
		}
	}
}

func dup(files []string) map[string][]string {
	filenameList := make(map[string][]string)
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "main: %v\n", err)
			continue
		}
		countLines(f, arg, filenameList)
		f.Close()
	}
	return filenameList
}
