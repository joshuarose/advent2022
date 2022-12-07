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
	TOTAL_SPACE              = 70000000
	REQUIRED_AVAILABLE_SPACE = 30000000
)

// Directory will be a graph, following a quasi linked list pattern
type Directory struct {
	Parent      *Directory
	Name        string
	Files       []*File
	Directories []*Directory
}

func (d *Directory) CalculateSize() int {
	totalSize := 0
	for _, f := range d.Files {
		totalSize += f.Size
	}
	for _, sd := range d.Directories {
		totalSize += sd.CalculateSize()
	}
	return totalSize
}

type File struct {
	Name string
	Size int
}

var fileName = "input.txt"
var rootDir *Directory
var currentDir *Directory
var sizes []int

func init() {
	rootDir = &Directory{
		Name:        "/",
		Files:       []*File{},
		Directories: []*Directory{},
		Parent:      nil,
	}
	currentDir = rootDir
	sizes = make([]int, 0)
}

func main() {
	fmt.Println(processCommandInput())
}

// ultimately returns the sum of all directories containing over 100000 elf bytes
func processCommandInput() int {
	hydrateFileGraph()
	processDirectorySize(rootDir)
	rightSize := rightSize()
	fmt.Printf("TOTAL USED SPACE: %d\n", totalUsedSpace())
	fmt.Printf("TOTAL AVAILABLE SPACE: %d\n", availableSpace())
	return rightSize
}

func hydrateFileGraph() {
	cmdFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Can't open txt file for command inputs")
	}

	cmdScanner := bufio.NewScanner(cmdFile)
	cmdScanner.Split(bufio.ScanLines)

	for cmdScanner.Scan() {
		processCommandLine(cmdScanner.Text())
	}
	cmdFile.Close()
}

func processCommandLine(command string) {
	if isParentDirectory(command) {
		currentDir = currentDir.Parent
		return
	}
	if isRootDirectory(command) {
		currentDir = rootDir
		return
	}
	if isChangeDirectory(command) {
		changeDirectory(command)
		return
	}
	if directoryListing(command) {
		addDirectory(command)
		return
	}
	if fileListing(command) {
		addFile(command)
	}
}

func isParentDirectory(command string) bool {
	if strings.Contains(command, "$ cd ..") {
		return true
	}
	return false
}

func isRootDirectory(command string) bool {
	if strings.Contains(command, "$ cd /") {
		return true
	}
	return false
}

func isChangeDirectory(command string) bool {
	if strings.Contains(command, "$ cd") && !isParentDirectory(command) && !isRootDirectory(command) {
		return true
	}
	return false
}

func directoryListing(command string) bool {
	if strings.Contains(command, "dir") {
		return true
	}
	return false
}

func fileListing(command string) bool {
	// defensively guard against the list command
	if command == "$ ls" {
		return false
	}
	// This assumes the only possibility left is file listing
	return true
}

func addDirectory(command string) {
	dirName := strings.ReplaceAll(command, "dir ", "")
	currentDir.Directories = append(currentDir.Directories, &Directory{
		Parent:      currentDir,
		Name:        dirName,
		Directories: []*Directory{},
	})
}

func addFile(command string) {
	cmdParts := strings.Split(command, " ")
	size, err := strconv.Atoi(cmdParts[0])
	if err != nil {
		return
	}
	currentDir.Files = append(currentDir.Files, &File{
		Name: cmdParts[1],
		Size: size,
	})
}

func changeDirectory(command string) {
	destDir := strings.ReplaceAll(command, "$ cd ", "")
	for _, v := range currentDir.Directories {
		if v.Name == destDir {
			currentDir = v
		}
	}
}

func processDirectorySize(dir *Directory) {
	total := dir.CalculateSize()
	sizes = append(sizes, total)
	for _, v := range dir.Directories {
		processDirectorySize(v)
	}
}

func totalUsedSpace() int {
	return rootDir.CalculateSize()
}

func availableSpace() int {
	return TOTAL_SPACE - totalUsedSpace()
}

func rightSize() int {
	rightSize := TOTAL_SPACE
	for _, v := range sizes {
		if v > (REQUIRED_AVAILABLE_SPACE - availableSpace()) {
			if v < rightSize {
				rightSize = v
			}
		}
	}
	return rightSize
}
