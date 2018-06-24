package main

import (
	"os"
	"testing"
)

func TestIsPresent(t *testing.T) {
	stringList := []string{"a", "b", "c", "d"}
	if result := isPresent(stringList, "a"); result != true {
		t.Errorf("Expected return val for 'a' is true, got %t", result)
	}

	if result := isPresent(stringList, "x"); result != false {
		t.Errorf("Expected return val for 'x' is false, got %t", result)
	}

	if result := isPresent([]string{}, "x"); result != false {
		t.Errorf("Expected return val for 'x' is false, got %t", result)
	}
}

func TestCountLines(t *testing.T) {
	answer := map[string][]string{
		"test": []string{"input_for_test1", "input_for_test2"},
		"get":  []string{"input_for_test1", "input_for_test2"},
		"aaa":  []string{"input_for_test1"},
		"bbb":  []string{"input_for_test1"},
		"bb":   []string{"input_for_test2"},
		"ddd":  []string{"input_for_test2"},
		"ccc":  []string{"input_for_test2"},
	}

	filenameList := make(map[string][]string)
	filename := "input_for_test1"
	f, err := os.Open(filename)
	if err != nil {
		t.Errorf("file open error: %v", err)
	}
	countLines(f, filename, filenameList)

	filename = "input_for_test2"
	f, err = os.Open(filename)
	if err != nil {
		t.Errorf("file open error: %v", err)
	}
	countLines(f, filename, filenameList)

	// line and files in answer is present in filenamelist
	for line, list := range answer {
		for _, filename := range list {
			if isPresent(filenameList[line], filename) != true {
				t.Errorf("Expected return val for %s is %v, got %v", line, list, filenameList[line])
			}
		}
	}

	// line and files in filenamelist is present in answer
	for line, list := range filenameList {
		for _, filename := range list {
			if isPresent(answer[line], filename) != true {
				t.Errorf("Expected return val for %s is %v, got %v", line, list, answer[line])
			}
		}
	}
}

func TestDup(t *testing.T) {
	files := []string{"input_for_test1", "input_for_test2"}
	answer := map[string][]string{
		"test": []string{"input_for_test1", "input_for_test2"},
		"get":  []string{"input_for_test1", "input_for_test2"},
		"aaa":  []string{"input_for_test1"},
		"bbb":  []string{"input_for_test1"},
		"bb":   []string{"input_for_test2"},
		"ddd":  []string{"input_for_test2"},
		"ccc":  []string{"input_for_test2"},
	}

	filenameList := dup(files)
	// line and files in answer is present in filenamelist
	for line, list := range answer {
		for _, filename := range list {
			if isPresent(filenameList[line], filename) != true {
				t.Errorf("Expected return val for %s is %v, got %v", line, list, filenameList[line])
			}
		}
	}

	// line and files in filenamelist is present in answer
	for line, list := range filenameList {
		for _, filename := range list {
			if isPresent(answer[line], filename) != true {
				t.Errorf("Expected return val for %s is %v, got %v", line, list, answer[line])
			}
		}
	}
}
