package main

import (
	"fmt"
	"log"
)

func max(vals ...float32) float32 {
	if len(vals) < 1 {
		log.Fatal("The length of the sequence is at least 1 or more.")
	}
	n := vals[0]
	for _, val := range vals {
		if n < val {
			n = val
		}
	}
	return n
}

func min(vals ...float32) float32 {
	if len(vals) < 1 {
		log.Fatal("The length of the sequence is at least 1 or more.")
	}
	n := vals[0]
	for _, val := range vals {
		if n > val {
			n = val
		}
	}
	return n
}

func maxWithOneOrMoreArg(v float32, vals ...float32) float32 {
	n := v
	for _, val := range vals {
		if n < val {
			n = val
		}
	}
	return n
}

func minWithOneOrMoreArg(v float32, vals ...float32) float32 {
	n := v
	for _, val := range vals {
		if n > val {
			n = val
		}
	}
	return n
}

func main() {
	values := []float32{1, 2, 3, 4}
	fmt.Println(max(values...))
	fmt.Println(min(values...))

	fmt.Println(maxWithOneOrMoreArg(values[0], values[1:]...))
	fmt.Println(minWithOneOrMoreArg(values[0], values[1:]...))
}
