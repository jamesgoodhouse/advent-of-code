package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	inputFile = "input.txt"
)

func main() {
	part1()
	part2()
}

func part1() {
	var increaseCounter, decreaseCounter int
	var previousValue *int

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// https://stackoverflow.com/a/16615559/12369692
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if previousValue != nil {
			if value > *previousValue {
				increaseCounter++
			} else if value < *previousValue {
				decreaseCounter++
			}
		}
		previousValue = &value
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part One")
	fmt.Printf("  Decreased: %d\n", decreaseCounter)
	fmt.Printf("  Increased: %d\n", increaseCounter)
}

func part2() {
	var increaseCounter, decreaseCounter int
	var previousValue *int
	var values []int

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// https://stackoverflow.com/a/16615559/12369692
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// https://www.geeksforgeeks.org/window-sliding-technique/
	for i := 0; i < len(values)-3+1; i++ {
		value := 0
		for j := 0; j < 3; j++ {
			value = value + values[i+j]
		}

		if previousValue != nil {
			if value > *previousValue {
				increaseCounter++
			} else if value < *previousValue {
				decreaseCounter++
			}
		}
		previousValue = &value
	}

	fmt.Println("Part Two")
	fmt.Printf("  Decreased: %d\n", decreaseCounter)
	fmt.Printf("  Increased: %d\n", increaseCounter)
}
