package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// var cycle int
var fileName = "input.txt"
var cmdLength = 144
var pw *PixelWriter
var commands []string
var rows [6]string

type PixelWriter struct {
	sync.Mutex
	headIndex int
}

func (v *PixelWriter) UpdateValue(modifier int) {
	v.headIndex += modifier
}

func init() {
	pw = &PixelWriter{
		headIndex: 1,
	}
	rows = [6]string{}
}

func main() {
	fmt.Println(processElfFile())
}

func processElfFile() int {
	commands = make([]string, cmdLength)
	cmdFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Can't open txt file for command inputs")
	}

	cmdScanner := bufio.NewScanner(cmdFile)
	cmdScanner.Split(bufio.ScanLines)
	for i := 0; i < cmdLength; i++ {
		if cmdScanner.Scan() {
			commands[i] = cmdScanner.Text()
		}
	}
	cmdFile.Close()
	getScreenOutput()
	totalPoints := 0
	for _, v := range rows {
		fmt.Println(v)
	}
	return totalPoints
}

func getScreenOutput() {
	cycle := 1
	currentRow := 0
	currentColumn := 0
	modifier := 0
	for _, cmd := range commands {
		iterator := 1
		if cmd != "noop" {
			iterator = 2
			modifier, _ = strconv.Atoi(strings.Split(cmd, " ")[1])
		}
		for i := 0; i < iterator; i++ {
			pixel := "."
			switch pw.headIndex {
			case currentColumn - 1:
				pixel = "#"
			case currentColumn:
				pixel = "#"
			case currentColumn + 1:
				pixel = "#"
			}

			rows[currentRow] += pixel
			cycle += 1
			currentColumn += 1

			if currentColumn > 39 {
				currentColumn = 0
				currentRow += 1
			}
		}
		pw.UpdateValue(modifier)
	}
}

// func enqueueCommand(command string) {
// 	if command == "noop" {
// 		return
// 	} else {
// 		cmdQueue = strings.ReplaceAll(command, "addx ", "")
// 	}
// }

// func processModifyRegister() {
// 	mod, _ := strconv.Atoi(cmdQueue)
// 	pw.UpdateValue(mod)
// 	cmdQueue = ""
// }

// func print(cycle int, sprite bool) {
// 	if cycle <= 240 {
// 		if !sprite {
// 			fmt.Print(" ")
// 		} else {
// 			fmt.Println("#")
// 		}
// 		switch cycle {
// 		case 40:
// 			fmt.Println()
// 		case 80:
// 			fmt.Println()
// 		case 120:
// 			fmt.Println()
// 		case 160:
// 			fmt.Println()
// 		case 200:
// 			fmt.Println()
// 		case 240:
// 			fmt.Println()
// 		}
// 	}
// }
