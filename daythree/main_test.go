package main

import "testing"

func TestGetItemPriority(t *testing.T) {
	cases := []struct {
		letter   string
		expected int
	}{
		{"a", 1},
		{"d", 4},
		{"z", 26},
		{"A", 27},
		{"B", 28},
		{"Z", 52},
	}

	for _, tc := range cases {
		result := getItemPriority(tc.letter)
		if result != tc.expected {
			t.Errorf("expected %d, but received %d", tc.expected, result)
		}
	}
}

func TestProcessLine(t *testing.T) {
	cases := []struct {
		line     string
		expected int
	}{
		{"aBcB", 28},
		{"abcd", 0},
	}
	for _, tc := range cases {
		result := processLine(tc.line)
		if result != tc.expected {
			t.Errorf("expected %d, but received %d", tc.expected, result)
		}
	}
}

func TestProcessElfFile(t *testing.T) {
	result := processElfFile()
	expected := 2567
	if result != expected {
		t.Errorf("expected %d, but received %d", expected, result)
	}
}

func TestGetBadgeType(t *testing.T) {
	cases := []struct {
		rucksacks []string
		expected  string
	}{
		{rucksacks: []string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
		}, expected: "r"},
		{rucksacks: []string{
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		}, expected: "Z"},
		{rucksacks: []string{
			"abc",
			"def",
		}, expected: ""},
	}

	for _, tc := range cases {
		result := getBadgeType(tc.rucksacks)
		if result != tc.expected {
			t.Errorf("expected: %s, but received %s", tc.expected, result)
		}
	}
}
