package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var gridSize = 99

var grid [][]int
var fileName = "input.txt"

func init() {
	grid = make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}
}

func main() {
	fmt.Println(processElfFile())
}

func processElfFile() int {
	gridFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Can't open txt file for command inputs")
	}

	cmdScanner := bufio.NewScanner(gridFile)
	cmdScanner.Split(bufio.ScanLines)
	rowNum := 0
	for cmdScanner.Scan() {
		hydrateGridRow(rowNum, cmdScanner.Text())
		rowNum += 1
	}
	gridFile.Close()
	visibleTreeCount := findTreeScore()
	return visibleTreeCount
}

func hydrateGridRow(rowNum int, row string) {
	for i, v := range row {
		treeHeight, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatalf("error parsing tree height in grid row %d column %d", rowNum, i)
		}
		grid[rowNum][i] = treeHeight
	}
}

func findTreeScore() int {
	highestTreeScore := 0
	// iterate over each grid row skipping the first + last row that are all visible
	for r := 0; r < gridSize; r++ {
		// iterate over each column in the row and examine height
		for c := 0; c < gridSize; c++ {
			currentScore := processTree(c, r)
			if currentScore > highestTreeScore {
				highestTreeScore = currentScore
			}
		}
	}
	return highestTreeScore
}

func processTree(c, r int) int {
	// look up
	up := lookUp(c, r)
	// look down
	down := lookDown(c, r)
	// look left
	left := lookLeft(c, r)
	// look right
	right := lookRight(c, r)
	return up * left * right * down
}

func lookUp(currentColumn, currentRow int) int {
	treeCount := 0
	for i := currentRow - 1; i >= 0; i-- {
		current := grid[currentRow][currentColumn]
		upDog := grid[i][currentColumn]
		if current <= upDog {
			treeCount += 1
			break
		}
		treeCount += 1
	}
	return treeCount
}

func lookDown(currentColumn, currentRow int) int {
	treeCount := 0
	for i := currentRow + 1; i < gridSize; i++ {
		current := grid[currentRow][currentColumn]
		downwardFacingDog := grid[i][currentColumn]
		if current <= downwardFacingDog {
			treeCount += 1
			break
		}
		treeCount += 1

	}
	return treeCount
}

func lookLeft(currentColumn, currentRow int) int {
	treeCount := 0
	for i := currentColumn - 1; i >= 0; i-- {
		current := grid[currentRow][currentColumn]
		leftShark := grid[currentRow][i]
		if current <= leftShark {
			treeCount += 1
			break
		}
		treeCount += 1
	}
	return treeCount
}

func lookRight(currentColumn, currentRow int) int {
	treeCount := 0
	for i := currentColumn + 1; i < gridSize; i++ {
		current := grid[currentRow][currentColumn]
		rightSaidFred := grid[currentRow][i]
		if current <= rightSaidFred {
			treeCount += 1
			break
		}
		treeCount += 1
	}
	return treeCount
}
