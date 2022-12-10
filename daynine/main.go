package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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

var tailVisitedPositions []Position
var knotPositions [10]Position
var fileName = "input.txt"

func init() {
	for i := 0; i < 10; i++ {
		knotPositions[i] = Position{0, 0}
	}
	tailVisitedPositions = append(tailVisitedPositions, Position{0, 0})
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
		knotPositions[0] = move(knotPositions[0], moveFunc)
		for i := 1; i < 10; i++ {
			moveTail(i)
			printGrid()
		}
	}
}

// The tail doesn't need direction it just follows the head always touching
func moveTail(knotIndex int) {
	// If the head is ever two steps directly up, down, left, or right from the tail,
	// the tail must also move one step in that direction so it remains close enough

	previousKnot := knotPositions[knotIndex-1]
	currentKnot := knotPositions[knotIndex]

	// Tail and Head are on same Y axis
	switch previousKnot.Y {
	case currentKnot.Y:
		if previousKnot.X-currentKnot.X > 1 {
			//Follow Right
			currentKnot = Position{currentKnot.X + 1, currentKnot.Y}
			addTailPosition(currentKnot, knotIndex)
		} else if previousKnot.X-currentKnot.X < -1 {
			// Follow Left
			currentKnot = Position{currentKnot.X - 1, currentKnot.Y}
			addTailPosition(currentKnot, knotIndex)
		}
		return
	}
	// Tail and Head are on same X axis
	switch previousKnot.X {
	case currentKnot.X:
		if previousKnot.Y-currentKnot.Y < -1 {
			// Follow Down
			currentKnot = Position{currentKnot.X, currentKnot.Y - 1}
			addTailPosition(currentKnot, knotIndex)
		} else if previousKnot.Y-currentKnot.Y > 1 {
			// Follow Up
			currentKnot = Position{currentKnot.X, currentKnot.Y + 1}
			addTailPosition(currentKnot, knotIndex)
		}
		return
	}
	if isTouching(previousKnot, currentKnot) {
		return
	}
	// Otherwise, if the head and tail aren't touching and aren't in the same row or column,
	// the tail always moves one step diagonally to keep up:
	if previousKnot.X-currentKnot.X >= 1 && previousKnot.Y-currentKnot.Y >= 2 {
		// move diagonally up and right H{4, 2} T{3, 0}
		currentKnot = Position{currentKnot.X + 1, currentKnot.Y + 1}
		addTailPosition(currentKnot, knotIndex)
	} else if previousKnot.X-currentKnot.X >= 2 && previousKnot.Y-currentKnot.Y >= 1 {
		// move diagonally up and right H{4, 2} T{3, 0}
		currentKnot = Position{currentKnot.X + 1, currentKnot.Y + 1}
		addTailPosition(currentKnot, knotIndex)
	} else if previousKnot.X-currentKnot.X <= -1 && previousKnot.Y-currentKnot.Y <= -2 {
		// move diagonally down and left H{3, 2} T{4, 4}
		currentKnot = Position{currentKnot.X - 1, currentKnot.Y - 1}
		addTailPosition(currentKnot, knotIndex)
	} else if previousKnot.X-currentKnot.X <= -2 && previousKnot.Y-currentKnot.Y <= -1 {
		// move diagonally down and left H{3, 2} T{4, 4}
		currentKnot = Position{currentKnot.X - 1, currentKnot.Y - 1}
		addTailPosition(currentKnot, knotIndex)
	} else if previousKnot.X-currentKnot.X >= 1 && previousKnot.Y-currentKnot.Y <= -2 {
		// move diagonally down and right
		currentKnot = Position{currentKnot.X + 1, currentKnot.Y - 1}
		addTailPosition(currentKnot, knotIndex)
	} else if previousKnot.X-currentKnot.X >= 2 && previousKnot.Y-currentKnot.Y <= -1 {
		// move diagonally down and right
		currentKnot = Position{currentKnot.X + 1, currentKnot.Y - 1}
		addTailPosition(currentKnot, knotIndex)
	} else if previousKnot.X-currentKnot.X <= -2 && previousKnot.Y-currentKnot.Y >= 1 {
		// move diagonally up and left
		currentKnot = Position{currentKnot.X - 1, currentKnot.Y + 1}
		addTailPosition(currentKnot, knotIndex)
	} else if previousKnot.X-currentKnot.X <= -1 && previousKnot.Y-currentKnot.Y <= 2 {
		// move diagonally up and left
		currentKnot = Position{currentKnot.X - 1, currentKnot.Y + 1}
		addTailPosition(currentKnot, knotIndex)
	}

}

func addTailPosition(p Position, index int) {
	knotPositions[index] = p
	if !slices.Contains(tailVisitedPositions, p) && index == 9 {
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

func isTouching(head, tail Position) bool {
	verticalDiff := head.Y - tail.Y
	horizontalDiff := head.X - tail.X
	if verticalDiff <= 1 && verticalDiff >= -1 && horizontalDiff <= 1 && horizontalDiff >= -1 {
		return true
	}
	return false
}

func printGrid() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Knot%d: x:%d y:%d\n", i, knotPositions[i].X, knotPositions[i].Y)
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Print("\033[H\033[2J")
}
