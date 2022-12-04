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

func main() {
	fmt.Println(processElfInput())
}

func processElfInput() int {
	txtfile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Can't open txt file for section assignments")
	}

	txtScanner := bufio.NewScanner(txtfile)
	txtScanner.Split(bufio.ScanLines)
	overlapCount := 0
	for txtScanner.Scan() {
		line := txtScanner.Text()
		pair := processLine(line)
		rangeOne := getSectionArray(pair[0])
		rangeTwo := getSectionArray(pair[1])
		if sectionsContainOverlap(rangeOne, rangeTwo) {
			overlapCount += 1
		}
	}
	txtfile.Close()
	return overlapCount
}

// Reading the description I think part 1 is focused on full overlapping pairs and not the total
// I'm saving this utility function for future in case part 2 expands that
func processLine(elfPair string) []string {
	return strings.Split(elfPair, ",")
}

// min section is 1 and max is 99

// takes 2-4 and returns integer array of sections
func getSectionArray(dashedSections string) []int {
	sectionArr := []int{}
	sections := strings.Split(dashedSections, "-")
	startingSection, err := strconv.Atoi(string(sections[0]))
	if err != nil {
		return sectionArr
	}
	endingSection, err := strconv.Atoi(string(sections[len(sections)-1]))
	if err != nil {
		return sectionArr
	}
	for startingSection <= endingSection {
		sectionArr = append(sectionArr, startingSection)
		startingSection += 1
	}
	return sectionArr
}

func sectionsFullyOverlap(pairOne, pairTwo []int) bool {
	if len(pairOne) < 1 || len(pairTwo) < 1 {
		return false
	}
	fullOverlap := true
	var smallerPair, largerPair []int
	if len(pairOne) > len(pairTwo) {
		smallerPair, largerPair = pairTwo, pairOne
	} else {
		smallerPair, largerPair = pairOne, pairTwo
	}
	for _, v := range smallerPair {
		if !slices.Contains(largerPair, v) {
			fullOverlap = false
		}
	}
	return fullOverlap
}

func sectionsContainOverlap(pairOne, pairTwo []int) bool {
	if len(pairOne) < 1 || len(pairTwo) < 1 {
		return false
	}
	partialOverlap := false
	var smallerPair, largerPair []int
	if len(pairOne) > len(pairTwo) {
		smallerPair, largerPair = pairTwo, pairOne
	} else {
		smallerPair, largerPair = pairOne, pairTwo
	}
	for _, v := range smallerPair {
		if slices.Contains(largerPair, v) {
			partialOverlap = true
		}
	}
	return partialOverlap
}
