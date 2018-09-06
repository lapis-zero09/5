package main

import (
	"fmt"
	"os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	c := os.Args[1]
	dist := map[string]int{}
	dist[c] = 0
	breadthFirst(func(item string) []string {
		for _, course := range prereqs[item] {
			if _, ok := dist[course]; ok {
				continue
			}
			dist[course] = dist[item] + 1
		}
		return prereqs[item]
	}, []string{c})

	for course, d := range dist {
		fmt.Println(course, ":", d)
	}
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
