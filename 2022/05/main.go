package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jamesgoodhouse/advent-of-code/2022/helper"
)

const (
	inputPath = "input.txt"

	stackWidth      = 3
	stackSpaceWidth = 1
)

var ()

func main() {
	numStacks, maxStackHeight := findStackDimensions(inputPath)

	stacks := loadStack(numStacks, maxStackHeight, inputPath)

	for _, row := range stacks {
		fmt.Printf("%+v\n", row)
	}
}

func findStackDimensions(inputPath string) (int, int) {
	scanner, inputFile, err := helper.NewFileScanner(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	maxStackHeight := 0
	numStacks := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, " 1") {
			numStacks = len(strings.Fields(line))
			break
		}
		maxStackHeight += 1
	}

	return numStacks, maxStackHeight
}

func loadStack(numStacks, maxHeight int, input string) [][]*string {
	stackSlice := make([][]*string, maxHeight)
	for i := range stackSlice {
		stackSlice[i] = make([]*string, numStacks)
	}

	scanner, inputFile, err := helper.NewFileScanner(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	lineNo := 0
	for scanner.Scan() {
		if lineNo == maxHeight {
			break
		}

		currentStack := 1
		currentStackStart := 1
		stackSpacing := 4

		stackRow := scanner.Text()

		stackRowSlice := strings.Split(stackRow, "")

		for currentStackStart < len(stackRowSlice) {
			stackValue := strings.Join(stackRowSlice[currentStackStart:currentStackStart+1], "")
			if strings.TrimSpace(stackValue) != "" {
				stackSlice[lineNo][currentStack-1] = &stackValue
			}
			currentStackStart += stackSpacing
			currentStack += 1
		}

		lineNo += 1
	}

	return stackSlice
}
