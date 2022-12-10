package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	RIGHT = "R"
	LEFT  = "L"
	UP    = "U"
	DOWN  = "D"
)

type Position struct {
	X int
	Y int
}

type Move struct {
	Direction string
	Spaces    int
}

var headVisitedPositions []Position
var tailVisitedPositions []Position
var startingPosition = Position{0, 0}
var headPosition = Position{0, 0}
var tailPosition = Position{0, 0}
var fileName = "input.txt"

func init() {
	headVisitedPositions = append(headVisitedPositions, startingPosition)
	tailVisitedPositions = append(tailVisitedPositions, startingPosition)
}

func main() {
	fmt.Println(processElfFile())
}

func processElfFile() int {
	moveFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Can't open txt file for command inputs")
	}

	moveScanner := bufio.NewScanner(moveFile)
	moveScanner.Split(bufio.ScanLines)
	rowNum := 0
	for moveScanner.Scan() {
		m := parseMovement(moveScanner.Text())
		moveHead(m.Direction, m.Spaces)
		rowNum += 1
	}
	moveFile.Close()
	return len(tailVisitedPositions)
}

func parseMovement(row string) Move {
	parts := strings.Split(row, " ")
	spaces, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return Move{
		Direction: parts[0],
		Spaces:    spaces,
	}
}

func moveHead(direction string, spaces int) {
	var moveFunc func(p Position) Position
	switch direction {
	case LEFT:
		moveFunc = moveLeft
	case RIGHT:
		moveFunc = moveRight
	case UP:
		moveFunc = moveUp
	case DOWN:
		moveFunc = moveDown
	}
	for i := 0; i < spaces; i++ {
		headPosition = move(headPosition, moveFunc)
		headVisitedPositions = append(headVisitedPositions, headPosition)
		moveTail()
	}
}

// The tail doesn't need direction it just follows the head always touching
func moveTail() {
	// If the head is ever two steps directly up, down, left, or right from the tail,
	// the tail must also move one step in that direction so it remains close enough

	// Tail and Head are on same Y axis
	if isTouching() {
		return
	}
	if headPosition.X == tailPosition.X && headPosition.Y == tailPosition.Y {
		return
	}
	switch headPosition.Y {
	case tailPosition.Y:
		if headPosition.X-tailPosition.X > 0 {
			//Follow Right
			tailPosition = Position{tailPosition.X + 1, tailPosition.Y}
			addTailPosition(tailPosition)
		} else if headPosition.X-tailPosition.X < 0 {
			// Follow Left
			tailPosition = Position{tailPosition.X - 1, tailPosition.Y}
			addTailPosition(tailPosition)
		}
		return
	}
	// Tail and Head are on same X axis
	switch headPosition.X {
	case tailPosition.X:
		if headPosition.Y-tailPosition.Y < 0 {
			// Follow Down
			tailPosition = Position{tailPosition.X, tailPosition.Y - 1}
			addTailPosition(tailPosition)
		} else if headPosition.Y-tailPosition.Y > 0 {
			// Follow Up
			tailPosition = Position{tailPosition.X + 1, tailPosition.Y}
			addTailPosition(tailPosition)
		}
		return
	}

	// Otherwise, if the head and tail aren't touching and aren't in the same row or column,
	// the tail always moves one step diagonally to keep up:
	allowedMoves := 2
	movedHorizontally := false
	for allowedMoves > 0 {
		if tailPosition.X < headPosition.X && !movedHorizontally {
			// move right
			tailPosition = Position{tailPosition.X + 1, tailPosition.Y}
			movedHorizontally = true
		} else if tailPosition.X > headPosition.X && !movedHorizontally {
			// move left
			tailPosition = Position{tailPosition.X - 1, tailPosition.Y}
			movedHorizontally = true
		} else if tailPosition.Y > headPosition.Y {
			// move down
			tailPosition = Position{tailPosition.X, tailPosition.Y - 1}
		} else if tailPosition.Y < headPosition.Y {
			// move up
			tailPosition = Position{tailPosition.X, tailPosition.Y + 1}
		}
		allowedMoves -= 1
	}
	addTailPosition(tailPosition)
	return
}

func addTailPosition(p Position) {
	if !slices.Contains(tailVisitedPositions, p) {
		tailVisitedPositions = append(tailVisitedPositions, p)
	}
}

func moveLeft(p Position) Position {
	return Position{p.X - 1, p.Y}
}

func moveRight(p Position) Position {
	return Position{p.X + 1, p.Y}
}

func moveUp(p Position) Position {
	return Position{p.X, p.Y + 1}
}

func moveDown(p Position) Position {
	return Position{p.X, p.Y - 1}
}

func move(p Position, f func(p Position) Position) Position {
	return f(p)
}

func isTouching() bool {
	verticalDiff := headPosition.Y - tailPosition.Y
	horizontalDiff := headPosition.X - tailPosition.X
	if verticalDiff <= 1 && verticalDiff >= -1 && horizontalDiff <= 1 && horizontalDiff >= -1 {
		return true
	}
	return false
}
