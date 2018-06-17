package main

import (
	"testing"
)

func TestEfficientEcho(t *testing.T) {
	s := []string{"test1", "test2", "あ　fds"}
	ret := efficientEcho(s)
	if answer := "test1 test2 あ　fds"; ret == answer {
		t.Errorf("Expected return val is %s, got %s", answer, ret)
	}
}

func TestInefficientEcho(t *testing.T) {
	s := []string{"test1", "test2", "あ　fds"}
	ret := inefficientEcho(s)
	if answer := "test1 test2 あ　fds"; ret == answer {
		t.Errorf("Expected return val is %s, got %s", answer, ret)
	}
}
