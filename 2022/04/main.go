package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jamesgoodhouse/advent-of-code/2022/helper"
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

	part1Count := 0

	for scanner.Scan() {
		line := scanner.Text()

		e1s, e1e, e2s, e2e := getElfSections(line)

		if e1s >= e2s &&
			e1e <= e2e {
			fmt.Printf("%v-%v contains %v-%v\n", e2s, e2e, e1s, e1e)
			part1Count += 1
		} else if e2s >= e1s &&
			e2e <= e1e {
			fmt.Printf("%v-%v contains %v-%v\n", e1s, e1e, e2s, e2e)
			part1Count += 1
		}
	}

	fmt.Println(part1Count)
}

func getElfSections(sections string) (int, int, int, int) {
	assignmentPairs := strings.Split(sections, ",")

	elf1Sections := strings.Split(assignmentPairs[0], "-")
	elf1SectionStart, _ := strconv.Atoi(elf1Sections[0])
	elf1SectionEnd, _ := strconv.Atoi(elf1Sections[1])

	elf2Sections := strings.Split(assignmentPairs[1], "-")
	elf2SectionStart, _ := strconv.Atoi(elf2Sections[0])
	elf2SectionEnd, _ := strconv.Atoi(elf2Sections[1])

	return elf1SectionStart,
		elf1SectionEnd,
		elf2SectionStart,
		elf2SectionEnd
}
