package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jamesgoodhouse/advent-of-code/2022/helper"
)

var (
	groupSackNum = 0
	groupSacks   = make([]string, 3)
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("no input file give")
		os.Exit(1)
	}

	scanner, file, err := helper.NewFileScanner(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var part1Sum rune
	var part2Sum rune

	for scanner.Scan() {
		line := scanner.Text()

		part1Sum += part1(line)

		groupSacks[groupSackNum] = line

		if groupSackNum == 2 {
			part2Sum += processGroupSacks()
			groupSackNum = 0
		} else {
			groupSackNum += 1
		}
	}

	fmt.Printf("Part 1: %v\n", part1Sum)
	fmt.Printf("Part 2: %v\n", part2Sum)
}

func part1(sack string) rune {
	firstCompartment := sack[0 : len(sack)/2]
	secondCompartment := sack[len(sack)/2:]

	for _, char := range firstCompartment {
		if strings.Contains(secondCompartment, string(char)) {
			return itemTypePriority(char)
			break
		}
	}

	return 0
}

func processGroupSacks() rune {
	for _, char := range groupSacks[0] {
		if strings.Contains(groupSacks[1], string(char)) {
			if strings.Contains(groupSacks[2], string(char)) {
				return itemTypePriority(char)
				break
			}
		}
	}

	return 0
}

func itemTypePriority(char rune) rune {
	if char >= 97 {
		// lowercase
		return char - 96
	} else {
		// uppercase
		return char - 38
	}

	return 0
}
