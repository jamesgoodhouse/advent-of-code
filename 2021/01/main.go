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

var (
	increaseCounter, decreaseCounter int
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var previousValue *int

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

	fmt.Printf("Decreased: %d\n", decreaseCounter)
	fmt.Printf("Increased: %d\n", increaseCounter)
}
