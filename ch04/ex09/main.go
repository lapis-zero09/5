package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cnts := make(map[string]int)

	fp, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	input := bufio.NewScanner(fp)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		cnts[input.Text()]++
	}

	fmt.Println("word\tcount")
	for word, count := range cnts {
		fmt.Printf("%s \t %d\n", word, count)
	}
}
