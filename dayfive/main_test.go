package main

import "testing"

func TestParseRow(t *testing.T) {
	cases := []struct {
		row      string
		expected []string
	}{
		{"        [H]     [W] [B]            ", []string{EMPTY_BOX, EMPTY_BOX, "[H]", EMPTY_BOX, "[W]", "[B]", EMPTY_BOX, EMPTY_BOX, EMPTY_BOX}},
		{"    [D] [B]     [L] [G] [N]        ", []string{EMPTY_BOX, "[D]", "[B]", EMPTY_BOX, "[L]", "[G]", "[N]", EMPTY_BOX, EMPTY_BOX}},
		{"[T] [L] [Z] [R] [C] [Q] [V] [P] [H]", []string{"[T]", "[L]", "[Z]", "[R]", "[C]", "[Q]", "[V]", "[P]", "[H]"}},
	}

	for _, tc := range cases {
		result := parseRow(tc.row)
		if len(result) != 9 {
			t.Errorf("Expected 9 stacks but received: %d", len(result))
		}
		for i, v := range result {
			if tc.expected[i] != v {
				t.Errorf("Expected array element at position %d to be %s", i, v)
			}
		}
	}
}

func TestParseInstruction(t *testing.T) {
	cases := []struct {
		row      string
		expected []int
	}{
		{"move 3 from 2 to 9", []int{3, 2, 9}},
		{"move 1 from 1 to 6", []int{1, 1, 6}},
		{"move 13 from 7 to 6", []int{13, 7, 6}},
	}
	for _, tc := range cases {
		result := parseInstruction(tc.row)
		if len(result) != len(tc.expected) {
			t.Error("Unexpected array length")
		}
		for i, v := range result {
			if tc.expected[i] != v {
				t.Errorf("expected %d but received %d", tc.expected[i], v)
			}
		}
	}
}

func TestSetupWarehouse(t *testing.T) {
	warehouseStrings := []string{
		"        [H]     [W] [B]            ",
		"    [D] [B]     [L] [G] [N]        ",
		"[P] [J] [T]     [M] [R] [D]        ",
		"[V] [F] [V]     [F] [Z] [B]     [C]",
		"[Z] [V] [S]     [G] [H] [C] [Q] [R]",
		"[W] [W] [L] [J] [B] [V] [P] [B] [Z]",
		"[D] [S] [M] [S] [Z] [W] [J] [T] [G]",
		"[T] [L] [Z] [R] [C] [Q] [V] [P] [H]",
	}
	setupWarehouse(warehouseStrings)
	result := warehouse[0][0]
	expected := "[T]"
	if result != expected {
		t.Errorf("mismatch warehouse: expected %s, received %s", expected, result)
	}
	result2 := len(warehouse[0])
	expected2 := 6
	if result2 != expected2 {
		t.Errorf("mismatch warehouse: expected %d, received %d", expected2, result2)
	}
}

func TestMoveCrates(t *testing.T) {
	warehouseStrings := []string{
		"        [H]     [W] [B]            ",
		"    [D] [B]     [L] [G] [N]        ",
		"[P] [J] [T]     [M] [R] [D]        ",
		"[V] [F] [V]     [F] [Z] [B]     [C]",
		"[Z] [V] [S]     [G] [H] [C] [Q] [R]",
		"[W] [W] [L] [J] [B] [V] [P] [B] [Z]",
		"[D] [S] [M] [S] [Z] [W] [J] [T] [G]",
		"[T] [L] [Z] [R] [C] [Q] [V] [P] [H]",
	}
	setupWarehouse(warehouseStrings)
	moveCrates([]int{3, 2, 9})
	lastRow := warehouse[8]
	expected := []string{"[H]", "[G]", "[Z]", "[R]", "[C]", "[D]", "[J]", "[F]"}
	for i, v := range expected {
		if lastRow[i] != v {
			t.Errorf("Expected element %s at index %d but received %s", expected[i], i, lastRow[i])
		}
	}
}

func TestProcessWarehouse(t *testing.T) {
	result := processWarehouse()
	expected := "TLFGBZHCN"
	if result != expected {
		t.Errorf("expected %s received %s", expected, result)
	}
}
