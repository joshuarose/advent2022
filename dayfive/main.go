package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	STACK_COUNT = 9
	EMPTY_BOX   = "   " // an empty box is 3 spaces in a slot
)

var warehouse [9][]string

func init() {
	warehouse = [9][]string{}
}

func main() {
	fmt.Println(processWarehouse())
}

// Sets up initial crate stacking from input.txt
func processWarehouse() string {
	// Open and process crates
	txtfile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Can't open txt file for section assignments")
	}

	txtScanner := bufio.NewScanner(txtfile)
	txtScanner.Split(bufio.ScanLines)
	currentRowIndex := 0
	fileRows := make([]string, 511)

	for txtScanner.Scan() {
		fileRows[currentRowIndex] = txtScanner.Text()
		currentRowIndex += 1
	}
	setupWarehouse(fileRows[0:8])
	for _, v := range fileRows[10:511] {
		instructions := parseInstruction(v)
		moveCrates(instructions)
	}
	crateMap, letters := friendlyPrintWarehouseData(warehouse)
	txtfile.Close()
	fmt.Println(crateMap)
	return letters
}

func parseRow(row string) []string {
	crates := make([]string, STACK_COUNT)
	spacers := []int{3, 7, 11, 15, 19, 23, 27, 31, 35}
	for i, v := range spacers {
		startCrate := v - 3
		crate := row[startCrate:v]
		crates[i] = crate
	}
	return crates
}

func parseInstruction(row string) []int {
	unfilteredArray := strings.Split(row, " ")
	moveQuantity, err := strconv.Atoi(unfilteredArray[1])
	if err != nil {
		log.Fatalln(err)
	}
	startingColumn, err := strconv.Atoi(unfilteredArray[3])
	if err != nil {
		log.Fatalln(err)
	}
	endingColumn, err := strconv.Atoi(unfilteredArray[5])
	if err != nil {
		log.Fatalln(err)
	}
	return []int{moveQuantity, startingColumn, endingColumn}
}

func moveCrates(instructions []int) {
	cratesToMove := instructions[0]
	fromStack := instructions[1] - 1
	toStack := instructions[2] - 1
	for i := 0; i < cratesToMove; i++ {
		topCrate := warehouse[fromStack][len(warehouse[fromStack])-1]
		warehouse[toStack] = append(warehouse[toStack], topCrate)
		warehouse[fromStack] = warehouse[fromStack][0 : len(warehouse[fromStack])-1]
	}
}

func setupWarehouse(rows []string) {
	// find tallest stack
	starterRows := [][]string{}
	for i := range rows {
		starterRows = append(starterRows, parseRow(rows[i]))
	}
	columns := [9][]string{}
	for i := 7; i >= 0; i-- {
		for j, v := range starterRows[i] {
			if v != EMPTY_BOX {
				columns[j] = append(columns[j], v)
			}
		}
	}
	warehouse = columns
}

func friendlyPrintWarehouseData(wh [9][]string) (string, string) {
	highestStack := 0
	topLetters := ""
	warehouseMap := ""
	for i := 0; i < 9; i++ {
		if len(wh[i]) > highestStack {
			highestStack = len(wh[i])
		}
		topLetters += wh[i][len(wh[i])-1]
	}
	for i := (highestStack); i > 0; i-- {
		for j := 0; j < 9; j++ {
			if len(wh[j]) >= i {
				warehouseMap += fmt.Sprint(wh[j][i-1])
			} else {
				warehouseMap += fmt.Sprint(EMPTY_BOX)
			}
		}
		warehouseMap += "\n"
	}
	topLetters = strings.ReplaceAll(topLetters, "[", "")
	topLetters = strings.ReplaceAll(topLetters, "]", "")
	return warehouseMap, topLetters
}
