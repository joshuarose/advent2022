package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func clearState() {
	headVisitedPositions = []Position{}
	tailVisitedPositions = []Position{}
	startingPosition = Position{0, 0}
	headPosition = Position{0, 0}
	tailPosition = Position{0, 0}
}

func TestProcessElfFile(t *testing.T) {
	result := processElfFile()
	// 14073 too high
	// 8017 too high
	// 7540 too high
	// 6801 ?
	// 6354
	expected := 6354
	if result != expected {
		t.Errorf("expected %d, received %d", expected, result)
	}
}

func TestLeft(t *testing.T) {
	clearState()
	result := moveLeft(startingPosition)
	expected := Position{-1, 0}
	if result.X != expected.X || result.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, result)
	}
}

func TestRight(t *testing.T) {
	clearState()
	result := moveRight(startingPosition)
	expected := Position{1, 0}
	if result.X != expected.X || result.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, result)
	}
}

func TestUp(t *testing.T) {
	clearState()
	result := moveUp(startingPosition)
	expected := Position{0, 1}
	if result.X != expected.X || result.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, result)
	}
}

func TestDown(t *testing.T) {
	clearState()
	result := moveDown(startingPosition)
	expected := Position{0, -1}
	if result.X != expected.X || result.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, result)
	}
}

func TestMove(t *testing.T) {
	clearState()
	result := move(startingPosition, moveDown)
	expected := Position{0, -1}
	if result.X != expected.X || result.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, result)
	}
}

func TestMoveHead(t *testing.T) {
	clearState()
	moveHead(UP, 3)
	expected := Position{0, 3}
	if headPosition.X != expected.X || headPosition.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, headPosition)
	}
	moveHead(RIGHT, 4)
	expected = Position{4, 3}
	if headPosition.X != expected.X || headPosition.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, headPosition)
	}
	moveHead(DOWN, 2)
	expected = Position{4, 1}
	if headPosition.X != expected.X || headPosition.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, headPosition)
	}
	moveHead(LEFT, 1)
	expected = Position{3, 1}
	if headPosition.X != expected.X || headPosition.Y != expected.Y {
		t.Errorf("expected %v but received %v", expected, headPosition)
	}
	expectedheadVisitedPositions := 10
	result := len(headVisitedPositions)
	if result != expectedheadVisitedPositions {
		t.Errorf("expected %d received %d", expectedheadVisitedPositions, result)
	}
}

func TestMoveTail(t *testing.T) {
	clearState()
	// .....    .....    .....
	// .TH.. -> .T.H. -> ..TH.
	// .....    .....    .....
	headPosition = Position{2, 1}
	tailPosition = Position{1, 1}
	moveHead(RIGHT, 1)
	if headPosition.X != 3 || headPosition.Y != 1 {
		t.Errorf("expected %v but received %v", Position{3, 1}, headPosition)
	}
	if tailPosition.X != 2 || tailPosition.Y != 1 {
		t.Errorf("expected %v but received %v", Position{2, 1}, tailPosition)
	}
	result := len(tailVisitedPositions)
	expected := 1
	if result != expected {
		t.Errorf("expected %d but received %d", expected, result)
	}
	// ...    ...    ...
	// .T.    .T.    ...
	// .H. -> ... -> .T.
	// ...    .H.    .H.
	// ...    ...    ...
	// Tail starts one above
	headPosition = Position{1, 2}
	tailPosition = Position{1, 3}
	moveHead(DOWN, 1)
	if headPosition.X != 1 || headPosition.Y != 1 {
		t.Errorf("expected %v but received %v", Position{1, 1}, headPosition)
	}
	if tailPosition.X != 1 || tailPosition.Y != 2 {
		t.Errorf("expected %v but received %v", Position{1, 2}, tailPosition)
	}
	result = len(tailVisitedPositions)
	expected = 2
	if result != expected {
		t.Errorf("expected %d but received %d", expected, result)
	}
	// .....    .....    .....
	// .....    ..H..    ..H..
	// ..H.. -> ..... -> ..T..
	// .T...    .T...    .....
	// .....    .....    .....
	headPosition = Position{2, 2}
	tailPosition = Position{1, 1}
	moveHead(UP, 1)
	if headPosition.X != 2 || headPosition.Y != 3 {
		t.Errorf("expected %v but received %v", Position{2, 3}, headPosition)
	}
	if tailPosition.X != 2 || tailPosition.Y != 2 {
		t.Errorf("expected %v but received %v", Position{2, 2}, tailPosition)
	}
	result = len(tailVisitedPositions)
	expected = 3
	if result != expected {
		t.Errorf("expected %d but received %d", expected, result)
	}
	// .....    .....    .....
	// .....    .....    .....
	// ..H.. -> ...H. -> ..TH.
	// .T...    .T...    .....
	// .....    .....    .....
	headPosition = Position{2, 2}
	tailPosition = Position{1, 1}
	moveHead(RIGHT, 1)
	if headPosition.X != 3 || headPosition.Y != 2 {
		t.Errorf("expected %v but received %v", Position{3, 2}, headPosition)
	}
	if tailPosition.X != 2 || tailPosition.Y != 2 {
		t.Errorf("expected %v but received %v", Position{2, 2}, tailPosition)
	}
	result = len(tailVisitedPositions)
	expected = 3
	if result != expected {
		t.Errorf("expected %d but received %d", expected, result)
	}
}

func TestAddTailPosition(t *testing.T) {
	clearState()
	expected := len(tailVisitedPositions)
	if expected != 0 {
		t.Fail()
	}
	addTailPosition(Position{0, 0})
	expected = len(tailVisitedPositions)
	if expected != 1 {
		t.Fail()
	}
	addTailPosition(Position{1, 1})
	expected = len(tailVisitedPositions)
	if expected != 2 {
		t.Fail()
	}
	addTailPosition(Position{1, 1})
	expected = len(tailVisitedPositions)
	if expected != 2 {
		t.Fail()
	}
}

func TestMovement(t *testing.T) {
	then := assert.New(t)
	// == Initial State ==
	// ......
	// ......
	// ......
	// ......
	// H.....  (H covers T, s)
	headPosition = Position{0, 0}
	tailPosition = Position{0, 0}

	// == R 4 ==
	moveHead(RIGHT, 4)

	// ......
	// ......
	// ......
	// ......
	// TH....  (T covers s)

	// ......
	// ......
	// ......
	// ......
	// sTH...

	// ......
	// ......
	// ......
	// ......
	// s.TH..

	// ......
	// ......
	// ......
	// ......
	// s..TH.
	then.Equal(tailPosition, Position{3, 0})
	then.Equal(headPosition, Position{4, 0})

	// == U 4 ==
	moveHead(UP, 4)

	// ......
	// ......
	// ......
	// ....H.
	// s..T..

	// ......
	// ......
	// ....H.
	// ....T.
	// s.....

	// ......
	// ....H.
	// ....T.
	// ......
	// s.....

	// ....H.
	// ....T.
	// ......
	// ......
	// s.....
	then.Equal(Position{4, 3}, tailPosition)
	then.Equal(Position{4, 4}, headPosition)

	// == L 3 ==
	moveHead(LEFT, 3)
	// ...H..
	// ....T.
	// ......
	// ......
	// s.....

	// ..HT..
	// ......
	// ......
	// ......
	// s.....

	// .HT...
	// ......
	// ......
	// ......
	// s.....
	then.Equal(Position{2, 4}, tailPosition)
	then.Equal(Position{1, 4}, headPosition)

	// == D 1 ==
	moveHead(DOWN, 1)

	// ..T...
	// .H....
	// ......
	// ......
	// s.....
	then.Equal(Position{2, 4}, tailPosition)
	then.Equal(Position{1, 3}, headPosition)
	// == R 4 ==
	moveHead(RIGHT, 4)
	// ..T...
	// ..H...
	// ......
	// ......
	// s.....

	// ..T...
	// ...H..
	// ......
	// ......
	// s.....

	// ......
	// ...TH.
	// ......
	// ......
	// s.....

	// ......
	// ....TH
	// ......
	// ......
	// s.....
	then.Equal(Position{4, 3}, tailPosition)
	then.Equal(Position{5, 3}, headPosition)

	// == D 1 ==
	moveHead(DOWN, 1)
	// ......
	// ....T.
	// .....H
	// ......
	// s.....
	then.Equal(Position{4, 3}, tailPosition)
	then.Equal(Position{5, 2}, headPosition)
	// == L 5 ==
	moveHead(LEFT, 5)
	// ......
	// ....T.
	// ....H.
	// ......
	// s.....

	// ......
	// ....T.
	// ...H..
	// ......
	// s.....

	// ......
	// ......
	// ..HT..
	// ......
	// s.....

	// ......
	// ......
	// .HT...
	// ......
	// s.....

	// ......
	// ......
	// HT....
	// ......
	// s.....
	then.Equal(Position{1, 2}, tailPosition)
	then.Equal(Position{0, 2}, headPosition)

	// == R 2 ==
	moveHead(RIGHT, 2)
	// ......
	// ......
	// .H....  (H covers T)
	// ......
	// s.....

	// ......
	// ......
	// .TH...
	// ......
	// s.....
	then.Equal(Position{1, 2}, tailPosition)
	then.Equal(Position{2, 2}, headPosition)
}
