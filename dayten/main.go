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

const (
	LISTEN_CYCLE_COUNT   = 5
	INITIAL_LISTEN_CYCLE = 20
	CYCLE_LISTEN_GAP     = 40
)

// var cycle int
var fileName = "input.txt"
var cmdLength = 144
var mainValue *ValueX
var commands []string
var cmdQueue string
var listenCycles [LISTEN_CYCLE_COUNT + 1]int
var signalPoints []int

type ValueX struct {
	sync.Mutex
	value int
}

func (v *ValueX) UpdateValue(modifier int) {
	v.value += modifier
}

func init() {
	listenCycles[0] = INITIAL_LISTEN_CYCLE
	listenCycles[1] = INITIAL_LISTEN_CYCLE + CYCLE_LISTEN_GAP
	for i := LISTEN_CYCLE_COUNT - 3; i <= LISTEN_CYCLE_COUNT; i++ {
		listenCycles[i] = listenCycles[i-1] + CYCLE_LISTEN_GAP
	}
	mainValue = &ValueX{
		value: 1,
	}
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
	commandIndex := 0
	for i := 1; i < 2000; i++ {
		showValue(i)
		if cmdQueue != "" {
			processModifyRegister()
		} else {
			if commandIndex < cmdLength-1 {
				enqueueCommand(commands[commandIndex])
				commandIndex += 1
			}
		}
	}
	totalPoints := 0
	for _, v := range signalPoints {
		totalPoints += v
	}
	return totalPoints
}

func enqueueCommand(command string) {
	if command == "noop" {
		return
	}
	cmdQueue = strings.ReplaceAll(command, "addx ", "")
}

func processModifyRegister() {
	mod, _ := strconv.Atoi(cmdQueue)
	mainValue.UpdateValue(mod)
	cmdQueue = ""
}

func showValue(cycle int) {
	switch cycle {
	case 20:
		fmt.Printf("CYCLE:%d VALUE:%d\n", cycle, mainValue.value*cycle)
		signalPoints = append(signalPoints, mainValue.value*cycle)
	case 60:
		fmt.Printf("CYCLE:%d VALUE:%d\n", cycle, mainValue.value*cycle)
		signalPoints = append(signalPoints, mainValue.value*cycle)
	case 100:
		fmt.Printf("CYCLE:%d VALUE:%d\n", cycle, mainValue.value*cycle)
		signalPoints = append(signalPoints, mainValue.value*cycle)
	case 140:
		fmt.Printf("CYCLE:%d VALUE:%d\n", cycle, mainValue.value*cycle)
		signalPoints = append(signalPoints, mainValue.value*cycle)
	case 180:
		fmt.Printf("CYCLE:%d VALUE:%d\n", cycle, mainValue.value*cycle)
		signalPoints = append(signalPoints, mainValue.value*cycle)
	case 220:
		fmt.Printf("CYCLE:%d VALUE:%d\n", cycle, mainValue.value*cycle)
		signalPoints = append(signalPoints, mainValue.value*cycle)
	}
}
