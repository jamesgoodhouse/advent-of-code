package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	inputFile  = "input.txt"
	numToSumTo = 2020
)

var (
	numbers = []int{}
)

func main() {
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

		var found bool

		for _, n1 := range numbers {
			for _, n2 := range numbers {
				if value+n1+n2 == numToSumTo {
					fmt.Println(value)
					fmt.Println(n1)
					fmt.Println(n2)
					fmt.Println(n1 * n2 * value)
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		numbers = append(numbers, value)
	}
}
