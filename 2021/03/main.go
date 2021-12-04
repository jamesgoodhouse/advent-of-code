package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ghetzel/go-stockutil/sliceutil"
)

const (
	inputFile = "input.txt"
)

func main() {
	var stringLen int
	var bitCount []int
	bitStrings := []string{}

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bitString := scanner.Text()
		bitStrings = append(bitStrings, bitString)

		if bitCount == nil {
			bitCount = make([]int, len(bitString))
			stringLen = len(bitString)
		}

		incrDecrBitCounts(bitString, &bitCount)
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

	oxStrings := bitStrings
	oxBitCount := bitCount
	co2Strings := bitStrings
	co2BitCount := bitCount

	for i := 0; i < stringLen; i++ {
		// already have counts for first iteration, so skip it
		if i > 0 {
			oxBitCount = make([]int, stringLen)
			for _, os := range oxStrings {
				incrDecrBitCounts(os, &oxBitCount)
			}

			co2BitCount = make([]int, stringLen)
			for _, cs := range co2Strings {
				incrDecrBitCounts(cs, &co2BitCount)
			}
		}

		ox := []string{}
		co2 := []string{}

		var oxValue string
		var co2Value string

		switch {
		case oxBitCount[i] > 0:
			oxValue = "1"
		case oxBitCount[i] < 0:
			oxValue = "0"
		case oxBitCount[i] == 0:
			oxValue = "1"
		}
		switch {
		case co2BitCount[i] > 0:
			co2Value = "0"
		case co2BitCount[i] < 0:
			co2Value = "1"
		case co2BitCount[i] == 0:
			co2Value = "0"
		}

		for k := range oxStrings {
			if oxStrings[k][i:i+1] == oxValue {
				if len(oxStrings) > 1 {
					ox = append(ox, oxStrings[k])
				}
			}
		}

		for k := range co2Strings {
			if co2Strings[k][i:i+1] == co2Value {
				if len(co2Strings) > 1 {
					co2 = append(co2, co2Strings[k])
				}
			}
		}

		if len(ox) > 0 {
			oxStrings = sliceutil.IntersectStrings(oxStrings, ox)
		}
		if len(co2) > 0 {
			co2Strings = sliceutil.IntersectStrings(co2Strings, co2)
		}
	}

	oxDec, err := strconv.ParseInt(oxStrings[0], 2, 64)
	co2Dec, err := strconv.ParseInt(co2Strings[0], 2, 64)

	fmt.Printf("\noxygen: %s (%d)\n", oxStrings[0], oxDec)
	fmt.Printf("co2: %s (%d)\n", co2Strings[0], co2Dec)
	fmt.Printf("multiplied: %d\n", oxDec*co2Dec)
}

func incrDecrBitCounts(s string, count *[]int) {
	for i := 0; i < len(s); i++ {
		switch s[i : i+1] {
		case "0":
			(*count)[i]--
		case "1":
			(*count)[i]++
		default:
			panic("something's wrong")
		}
	}
}
