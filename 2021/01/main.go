package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	inputFile  = "input.txt"
	windowSize = 3
)

func main() {
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
	for i := 0; i < len(values)-windowSize+1; i++ {
		value := 0
		for j := 0; j < windowSize; j++ {
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

	fmt.Printf("Decreased: %d\n", decreaseCounter)
	fmt.Printf("Increased: %d\n", increaseCounter)
}
