package main

import (
	"testing"
)

func TestTotalizeElfCalories(t *testing.T) {
	cases := []struct {
		calories []int
		record   []int
		expected []int
	}{
		{calories: []int{1000, 500, 200}, record: []int{1100, 300, 200}, expected: []int{1700, 1100, 300}},
		{calories: []int{1000, 500, 200, 100}, record: []int{2000, 1700, 1600}, expected: []int{2000, 1800, 1700}},
		{calories: []int{1000}, record: []int{1100, 1050, 900}, expected: []int{1100, 1050, 1000}},
		{calories: []int{1000}, record: []int{1100, 1050, 1010}, expected: []int{1100, 1050, 1010}},
	}
	for _, tc := range cases {
		_, records := totalizeElfCalories(tc.calories, tc.record)
		if records[0] != tc.expected[0] {
			t.Errorf("expected: %d received: %d", records[0], tc.expected[0])
		}
		if records[1] != tc.expected[1] {
			t.Errorf("expected: %d received: %d", records[1], tc.expected[1])
		}
		if records[2] != tc.expected[2] {
			t.Errorf("expected: %d received: %d", records[2], tc.expected[2])
		}
	}
}

func TestAddElfCalories(t *testing.T) {
	numbers := addElfCalories("10001", []int{1, 1, 2, 3})
	if numbers[len(numbers)-1] != 10001 {
		t.Error("Error appending number")
	}
}

func TestCalculateTotals(t *testing.T) {
	total := calculateTotals()
	expected := 205370
	if total != expected {
		t.Errorf("Incorrect total, expected %d, received %d", expected, total)
	}
}
