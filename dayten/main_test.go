package main

import (
	"testing"
)

func TestProcessFile(t *testing.T) {
	fileName = "test.txt"
	cmdLength = 146
	result := processElfFile()
	// 10160
	expected := 13140
	if result != expected {
		t.Errorf("expected %d received %d", expected, result)
	}
}
