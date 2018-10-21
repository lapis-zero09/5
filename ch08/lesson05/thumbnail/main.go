package main

import (
	"bufio"
	"log"
	"os"

	"github.com/adonovan/gopl.io/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
