package main

import "testing"

func TestMain(t *testing.T) {
	if 1+1 != 2 {
		t.Fail()
	}
}

func TestFindMarker(t *testing.T) {
	cases := []struct {
		transmission string
		expected     int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}
	for _, tc := range cases {
		result := findMarker(tc.transmission)
		if tc.expected != result {
			t.Errorf("expected: %d but received %d", tc.expected, result)
		}
	}
}
