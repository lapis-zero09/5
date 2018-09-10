package main

import (
	"bufio"
	"fmt"
	"log"
)

type Counter struct {
	wordsCount, linesCount int
}

func (c *Counter) WordsSize() int {
	return c.wordsCount
}

func (c *Counter) LinesSize() int {
	return c.linesCount
}

func (c *Counter) Write(p []byte) (int, error) {
	for b := p; len(b) > 0; {
		advance, _, err := bufio.ScanLines(b, true)
		if err != nil {
			log.Fatal(err)
		}
		b = b[advance:]
		c.linesCount++
	}

	for b := p; len(b) > 0; {
		advance, _, err := bufio.ScanWords(b, true)
		if err != nil {
			log.Fatal(err)
		}
		b = b[advance:]
		c.wordsCount++
	}
	return len(p), nil
}

func main() {
	s := "This is test sentence.\naaa\nbbb\n\nccc"
	var c Counter
	c.Write([]byte(s))
	fmt.Println(c)
}
