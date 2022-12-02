package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName = "input.txt"

// A for Rock, B for Paper, and C for Scissors - elf
// X for Rock, Y for Paper, and Z for Scissors - player
const (
	ROCK_SELECTED_POINTS     = 1
	PAPER_SELECTED_POINTS    = 2
	SCISSORS_SELECTED_POINTS = 3
	ROUND_WON_POINTS         = 6
	ROUND_DRAW_POINTS        = 3
	ROUND_LOST_POINTS        = 0
	ELF_PLAYS_ROCK           = "A"
	ELF_PLAYS_PAPER          = "B"
	ELF_PLAYS_SCISSORS       = "C"
	PLAYER_MUST_LOSE         = "X"
	PLAYER_MUST_DRAW         = "Y"
	PLAYER_MUST_WIN          = "Z"
)

func main() {
	fmt.Println(processElfFile())
}

// processes puzzle input and returns player score
func processElfFile() int {
	txtfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Can't open txt file for rock,paper,scissors inputs")
	}

	txtScanner := bufio.NewScanner(txtfile)
	txtScanner.Split(bufio.ScanLines)
	scores := []int{}
	for txtScanner.Scan() {
		roundPoints := processRound(txtScanner.Text())
		scores = append(scores, roundPoints)
	}
	txtfile.Close()
	totalPoints := 0
	for _, v := range scores {
		totalPoints += v
	}
	return totalPoints
}

func processRound(round string) int {
	plays := strings.Split(round, " ")
	if len(plays) < 2 {
		return 0
	}
	elfPlay := plays[0]
	playerPlay := plays[1]
	switch elfPlay {
	case ELF_PLAYS_ROCK:
		if playerPlay == PLAYER_MUST_WIN {
			return PAPER_SELECTED_POINTS + ROUND_WON_POINTS
		} else if playerPlay == PLAYER_MUST_DRAW {
			return ROCK_SELECTED_POINTS + ROUND_DRAW_POINTS
		} else {
			return SCISSORS_SELECTED_POINTS + ROUND_LOST_POINTS
		}
	case ELF_PLAYS_PAPER:
		if playerPlay == PLAYER_MUST_WIN {
			return SCISSORS_SELECTED_POINTS + ROUND_WON_POINTS
		} else if playerPlay == PLAYER_MUST_DRAW {
			return PAPER_SELECTED_POINTS + ROUND_DRAW_POINTS
		} else {
			return ROCK_SELECTED_POINTS + ROUND_LOST_POINTS
		}
	case ELF_PLAYS_SCISSORS:
		if playerPlay == PLAYER_MUST_WIN {
			return ROCK_SELECTED_POINTS + ROUND_WON_POINTS
		} else if playerPlay == PLAYER_MUST_DRAW {
			return SCISSORS_SELECTED_POINTS + ROUND_DRAW_POINTS
		} else {
			return PAPER_SELECTED_POINTS + ROUND_LOST_POINTS
		}
	}
	// invalid shouldn't hit here unless input has bad data
	return 0
}
