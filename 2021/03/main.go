package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ()

const (
	inputFile = "input.txt"
)

func main() {
	bitCount := make([]int, 12)

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bitString := scanner.Text()
		for i := 0; i < len(bitString); i++ {
			switch bitString[i : i+1] {
			case "0":
				bitCount[i]--
			case "1":
				bitCount[i]++
			default:
				panic("something's wrong")
			}
		}
	}

	var gamma string
	var epsilon string

	for i := range bitCount {
		switch {
		case bitCount[i] < 0:
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		case bitCount[i] > 0:
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		default:
			panic("something's wrong")
		}
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, err := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("gamma: %s (%d)\n", gamma, gammaDec)
	fmt.Printf("epsilon: %s (%d)\n", epsilon, epsilonDec)
	fmt.Printf("multiplied together: %d\n", epsilonDec*gammaDec)
}
