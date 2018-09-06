package main

import (
	"fmt"
	"log"
)

var prereqs = map[string]map[string]bool{
	"algorithms":     {"data structures": true},
	"calculus":       {"linear algebra": true},
	"linear algebra": {"calculus": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	topo, ok := topoSort(prereqs)
	if !ok {
		log.Fatal("topo loop!")
	}
	for i, course := range topo {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) ([]string, bool) {
	var order []string
	var loopFlag bool
	seen := make(map[string]int)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item := range items {
			switch seen[item] {
			case 0:
				seen[item]++
				visitAll(m[item])
				order = append(order, item)
				seen[item]++
			case 1:
				loopFlag = true
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)
	return order, !loopFlag
}
