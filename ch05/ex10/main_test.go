package main

import "testing"

func TestTopoSort(t *testing.T) {
	order := make(map[string]int)
	for idx, course := range topoSort(prereqs) {
		order[course] = idx
	}

	for course, deps := range prereqs {
		for dep := range deps {
			if order[course] < order[dep] {
				t.Errorf("unexpected result. expected: %v >= %v", order[course], order[dep])
			}
		}
	}
}
