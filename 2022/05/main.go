package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/jamesgoodhouse/advent-of-code/2022/helper"
)

const (
	inputPath    = "input.txt"
	stackSpacing = 4
)

var ()

func main() {
	numStacks, maxStackHeight := findStackDimensions(inputPath)

	stacks := loadStack(numStacks, maxStackHeight, inputPath)

	iLikeToMoveItMoveIt(stacks)

	spew.Dump(stacks)
}

func iLikeToMoveItMoveIt(stacks [][]*string) {
	pattern := regexp.MustCompile(`move\s(?P<move>\d+)\sfrom\s(?P<from>\d+)\sto\s(?P<to>\d+)`)

	scanner, inputFile, err := helper.NewFileScanner(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, "move") {
			continue
		}

		// https://gist.github.com/eculver/d1338aa87e87890e05d4f61ed0a33d6e
		matches := pattern.FindStringSubmatch(line)
		subnames := pattern.SubexpNames()
		if matches == nil || len(matches) != len(subnames) {
			panic("huh?")
		}

		instructionsMap := map[string]int{}
		for i := 1; i < len(matches); i++ {
			val, err := strconv.Atoi(matches[i])
			if err != nil {
				panic("uhh...")
			}
			instructionsMap[subnames[i]] = val
		}

		for i := 0; i < instructionsMap["move"]; i++ {
			val, newStack := stacks[instructionsMap["from"]-1][0], stacks[instructionsMap["from"]-1][1:]
			stacks[instructionsMap["from"]-1] = newStack
			stacks[instructionsMap["to"]-1] = append([]*string{val}, stacks[instructionsMap["to"]-1]...)
		}
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
	stackSlice := make([][]*string, numStacks)
	for i := range stackSlice {
		stackSlice[i] = make([]*string, maxHeight)
	}

	scanner, inputFile, err := helper.NewFileScanner(input)
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

		stackRow := scanner.Text()

		stackRowSlice := strings.Split(stackRow, "")

		for currentStackStart < len(stackRowSlice) {
			stackValue := strings.Join(stackRowSlice[currentStackStart:currentStackStart+1], "")
			if strings.TrimSpace(stackValue) != "" {
				stackSlice[currentStack-1][lineNo] = &stackValue
			}
			currentStackStart += stackSpacing
			currentStack += 1
		}

		lineNo += 1
	}

	for stackNum, stack := range stackSlice {
		newStack := []*string{}
		for _, item := range stack {
			if item != nil {
				newStack = append(newStack, item)
			}
		}
		stackSlice[stackNum] = newStack
	}

	return stackSlice
}
