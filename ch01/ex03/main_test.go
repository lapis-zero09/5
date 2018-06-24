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

func BenchmarkInefficientEcho(b *testing.B) {
	s := []string{"test1", "test2", "あ　fds"}
	for i := 0; i < b.N; i++ {
		inefficientEcho(s)
	}
}

func BenchmarkEfficientEcho(b *testing.B) {
	s := []string{"test1", "test2", "あ　fds"}
	for i := 0; i < b.N; i++ {
		efficientEcho(s)
	}
}
