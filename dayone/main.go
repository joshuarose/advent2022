package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	calorieRecord := calculateTotals()
	fmt.Println(calorieRecord)
}

func calculateTotals() int {
	txtfile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Can't open CSV file for calorie inputs")
	}

	txtScanner := bufio.NewScanner(txtfile)
	txtScanner.Split(bufio.ScanLines)

	calorieRecords := []int{0, 0, 0}
	currentElfCollection := []int{}

	for txtScanner.Scan() {
		value := txtScanner.Text()
		if value == "" {
			currentElfCollection, calorieRecords = totalizeElfCalories(currentElfCollection, calorieRecords)
		} else {
			currentElfCollection = addElfCalories(value, currentElfCollection)
		}
	}
	txtfile.Close()
	total := calorieRecords[0] + calorieRecords[1] + calorieRecords[2]
	return total
}

func totalizeElfCalories(elfCalories []int, calorieRecords []int) ([]int, []int) {
	if len(elfCalories) > 0 {
		currentTotal := 0
		for _, v := range elfCalories {
			currentTotal += v
		}
		elfCalories = []int{}
		firstPlace := calorieRecords[0]
		secondPlace := calorieRecords[1]
		thirdPlace := calorieRecords[2]
		if currentTotal > calorieRecords[0] {
			thirdPlace = secondPlace
			secondPlace = firstPlace
			firstPlace = currentTotal
		} else if currentTotal > calorieRecords[1] {
			thirdPlace = secondPlace
			secondPlace = currentTotal
		} else if currentTotal > calorieRecords[2] {
			thirdPlace = currentTotal
		}
		calorieRecords = []int{firstPlace, secondPlace, thirdPlace}
	}
	return elfCalories, calorieRecords
}

func addElfCalories(calories string, calorieArray []int) []int {
	calInt, err := strconv.Atoi(calories)
	if err != nil {
		log.Fatalln(err)
	}
	return append(calorieArray, calInt)
}
