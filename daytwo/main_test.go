package main

import (
	"fmt"
	"testing"
)

func TestProcessElfFile(t *testing.T) {
	result := processElfFile()
	expected := 12772
	if result != expected {
		t.Errorf("final score expected: %d, received: %d", expected, result)
	}
}

func TestProcessRounds(t *testing.T) {
	cases := []struct {
		round    string
		expected int
	}{
		{fmt.Sprintf("%s %s", ELF_PLAYS_ROCK, PLAYER_PLAYS_ROCK), ROCK_SELECTED_POINTS + ROUND_DRAW_POINTS},
		{fmt.Sprintf("%s %s", ELF_PLAYS_ROCK, PLAYER_PLAYS_PAPER), PAPER_SELECTED_POINTS + ROUND_WON_POINTS},
		{fmt.Sprintf("%s %s", ELF_PLAYS_ROCK, PLAYER_PLAYS_SCISSORS), SCISSORS_SELECTED_POINTS + ROUND_LOST_POINTS},
		{fmt.Sprintf("%s %s", ELF_PLAYS_PAPER, PLAYER_PLAYS_ROCK), ROCK_SELECTED_POINTS + ROUND_LOST_POINTS},
		{fmt.Sprintf("%s %s", ELF_PLAYS_PAPER, PLAYER_PLAYS_PAPER), PAPER_SELECTED_POINTS + ROUND_DRAW_POINTS},
		{fmt.Sprintf("%s %s", ELF_PLAYS_PAPER, PLAYER_PLAYS_SCISSORS), SCISSORS_SELECTED_POINTS + ROUND_WON_POINTS},
		{fmt.Sprintf("%s %s", ELF_PLAYS_SCISSORS, PLAYER_PLAYS_ROCK), ROCK_SELECTED_POINTS + ROUND_WON_POINTS},
		{fmt.Sprintf("%s %s", ELF_PLAYS_SCISSORS, PLAYER_PLAYS_PAPER), PAPER_SELECTED_POINTS + ROUND_LOST_POINTS},
		{fmt.Sprintf("%s %s", ELF_PLAYS_SCISSORS, PLAYER_PLAYS_SCISSORS), SCISSORS_SELECTED_POINTS + ROUND_DRAW_POINTS},
		{"INVALID", 0},
		{"T T", 0},
	}
	for _, tc := range cases {
		result := processRound(tc.round)
		if result != tc.expected {
			t.Errorf("Round: %s - Points expected: %d, received: %d", tc.round, tc.expected, result)
		}
	}
}
