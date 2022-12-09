package main

import "testing"

func TestProcessElfFile(t *testing.T) {
	result := processElfFile()
	expected := 440640
	if result != expected {
		t.Errorf("expected %d received %d", expected, result)
	}
}

func TestProcessTree(t *testing.T) {
	gridSize = 5
	grid = [][]int{[]int{3, 0, 3, 7, 3}, []int{2, 5, 5, 1, 2}, []int{6, 5, 3, 3, 2}, []int{3, 3, 5, 4, 9}, []int{3, 5, 3, 9, 0}}
	result := processTree(2, 1)
	expected := 4
	if result != expected {
		t.Errorf("expected %d received %d", expected, result)
	}
	result = processTree(2, 3)
	expected = 8
	if result != expected {
		t.Errorf("expected %d received %d", expected, result)
	}
}
