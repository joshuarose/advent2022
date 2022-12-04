package main

import "testing"

func TestMain(t *testing.T) {
	if 2+2 != 4 {
		t.Error("placeholder for day four")
	}
}

func TestGetSectionArray(t *testing.T) {
	cases := []struct {
		dashedSections string
		expected       []int
	}{
		{"2-4", []int{2, 3, 4}},
		{"6-8", []int{6, 7, 8}},
		{"2-3", []int{2, 3}},
		{"4-5", []int{4, 5}},
		{"5-7", []int{5, 6, 7}},
		{"7-9", []int{7, 8, 9}},
		{"6-6", []int{6}},
		{"2-8", []int{2, 3, 4, 5, 6, 7, 8}},
		{"7-X", []int{}},
		{"Y-7", []int{}},
		{"10-19", []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}},
	}
	for _, tc := range cases {
		result := getSectionArray(tc.dashedSections)
		if len(result) < len(tc.expected) {
			t.Error("Array size mismatch")
		}
		for i, v := range result {
			if tc.expected[i] != v {
				t.Errorf("Expected element in array: %d but received %d", tc.expected[i], v)
			}
		}
	}
}

func TestProcessLine(t *testing.T) {
	result := processLine("8-18,10-19")
	expected := []string{"8-18", "10-19"}
	if len(result) != 2 {
		t.Error("Expected resulting array to have 2 elements")
	}
	if result[0] != expected[0] && result[1] != expected[1] {
		t.Error("Expected matching arrays")
	}
}

func TestSectionsFullyOverlap(t *testing.T) {
	cases := []struct {
		sectionOne []int
		sectionTwo []int
		expected   bool
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, false},
		{[]int{4, 5, 6}, []int{1, 2, 3}, false},
		{[]int{2, 3, 4}, []int{1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4}, true},
		{[]int{6, 7, 8, 9, 10}, []int{8}, true},
		{[]int{99, 98, 97}, []int{97, 98, 99}, true},
		{[]int{}, []int{}, false},
		{[]int{1}, []int{1}, true},
		{[]int{9}, []int{1}, false},
		{[]int{6}, []int{4, 5, 6}, true},
		{[]int{2, 3, 4, 5, 6, 7, 8}, []int{3, 4, 5, 6, 7}, true},
		// 3-8,13-98
	}
	for _, tc := range cases {
		result := sectionsFullyOverlap(tc.sectionOne, tc.sectionTwo)
		if result != tc.expected {
			t.Errorf("Expected %v but received %v on pairs %v and %v", tc.expected, result, tc.sectionOne, tc.sectionTwo)
		}
	}
}

func TestSectionsPartiallyOverlap(t *testing.T) {
	cases := []struct {
		sectionOne []int
		sectionTwo []int
		expected   bool
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, false},
		{[]int{4, 5, 6}, []int{1, 2, 3}, false},
		{[]int{2, 3, 4}, []int{1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4}, true},
		{[]int{6, 7, 8, 9, 10}, []int{8}, true},
		{[]int{99, 98, 97}, []int{97, 98, 99}, true},
		{[]int{}, []int{}, false},
		{[]int{1}, []int{1}, true},
		{[]int{9}, []int{1}, false},
		{[]int{6}, []int{4, 5, 6}, true},
		{[]int{10, 11, 12, 13, 14, 15, 16}, []int{15, 16, 17}, true},
	}
	for _, tc := range cases {
		result := sectionsContainOverlap(tc.sectionOne, tc.sectionTwo)
		if result != tc.expected {
			t.Errorf("Expected %v but received %v on pairs %v and %v", tc.expected, result, tc.sectionOne, tc.sectionTwo)
		}
	}
}

func TestProcessElfInput(t *testing.T) {
	result := processElfInput()
	if result != 794 {
		t.Error("fail")
	}
}
