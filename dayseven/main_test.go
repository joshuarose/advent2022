package main

import "testing"

func TestProcessCommandInput(t *testing.T) {
	result := processCommandInput()
	expected := 1544176
	if result != expected {
		t.Errorf("expected %d but received %d", expected, result)
	}
}

func TestTotalUsedSpace(t *testing.T) {

	hydrateFileGraph()
	result := totalUsedSpace()
	expected := 83037906
	if result != expected {
		t.Errorf("expected %d but received %d", expected, result)
	}
}
