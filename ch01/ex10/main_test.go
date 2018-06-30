package main

import (
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	ch := make(chan string)
	go fetch("http://gopl.io", ch)
	fetchResult := strings.Split(<-ch, " ")

	if res := fetchResult[1]; res != "4154" {
		t.Errorf("Expected return is '4154', got %s", res)
	}

	if res := fetchResult[2]; res != "http://gopl.io" {
		t.Errorf("Expected return is 'http://gopl.io', got %s", res)
	}

}
