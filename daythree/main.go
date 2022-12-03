package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName = "input.txt"

func main() {
	fmt.Println(processElfFile())
}

// processes puzzle input and returns item sum
func processElfFile() int {
	txtfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Can't open txt file for rucksack inputs")
	}

	txtScanner := bufio.NewScanner(txtfile)
	txtScanner.Split(bufio.ScanLines)
	total := 0
	for txtScanner.Scan() {
		line := txtScanner.Text()
		total += processLine(line)
	}
	txtfile.Close()
	return total
}

// Process 1 line at a time from input.txt
func processLine(line string) int {
	lineLength := len(line)
	compartmentSize := lineLength / 2
	compartmentOne := ""
	compartmentTwo := ""
	var dupes []string
	for i := range line {
		if i < compartmentSize {
			compartmentOne += string(line[i])
		} else {
			compartmentTwo += string(line[i])
		}
	}
	for _, v := range compartmentOne {
		letter := string(v)
		if strings.Contains(compartmentTwo, letter) {
			dupes = append(dupes, letter)
		}
	}
	priorityTotal := 0
	for _, v := range dupes {
		priorityTotal = getItemPriority(v)
	}
	return priorityTotal
}

// Items are represented by alpha, upper and lower are different
// Priority value of a-z is 1-26, A-Z is 27-52
func getItemPriority(letter string) int {
	var lower = "abcdefghijklmnopqrstuvwxyz"
	var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	index := strings.Index(lower, letter)
	index = index + 1
	if index < 1 {
		index = strings.Index(upper, letter)
		index = index + 27
	}
	return index
}
