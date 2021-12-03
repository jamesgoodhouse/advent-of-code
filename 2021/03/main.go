package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ghetzel/go-stockutil/sliceutil"
)

const (
	inputFile = "input.txt"
)

// GROSSEST CODE EVER WRITTEN.

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
			log.Println("initializing bitCount")
			bitCount = make([]int, len(bitString))
			stringLen = len(bitString)
		}

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

	oxStrings := bitStrings
	oxBitCount := bitCount

	for i := 0; i < stringLen; i++ {
		log.Println("oxStrings is equal to...")
		log.Println(oxStrings)

		if i > 0 {
			oxBitCount = nil

			for _, os := range oxStrings {
				if oxBitCount == nil {
					oxBitCount = make([]int, len(os))
				}

				for j := 0; j < len(os); j++ {
					switch os[j : j+1] {
					case "0":
						oxBitCount[j]--
					case "1":
						oxBitCount[j]++
					default:
						panic("something's wrong")
					}
				}
			}
		}

		ox := []string{}

		switch {
		case oxBitCount[i] > 0:
			for k := range oxStrings {
				if oxStrings[k][i:i+1] == "1" {
					if len(oxStrings) > 1 {
						ox = append(ox, oxStrings[k])
					}
				}
			}
		case oxBitCount[i] < 0:
			for k := range oxStrings {
				if oxStrings[k][i:i+1] == "0" {
					if len(oxStrings) > 1 {
						ox = append(ox, oxStrings[k])
					}
				}
			}
		case oxBitCount[i] == 0:
			for k := range oxStrings {
				if oxStrings[k][i:i+1] == "1" {
					if len(oxStrings) > 1 {
						ox = append(ox, oxStrings[k])
					}
				}
			}
		}

		log.Println("ox is equal to...")
		log.Println(ox)

		if len(ox) > 0 {
			oxStrings = sliceutil.IntersectStrings(oxStrings, ox)
		}

		log.Println("oxStrings is now equal to...")
		log.Println(oxStrings)
		log.Println("")
	}

	co2Strings := bitStrings
	co2BitCount := bitCount

	for i := 0; i < stringLen; i++ {
		log.Println("co2Strings is equal to...")
		log.Println(co2Strings)

		if i > 0 {
			co2BitCount = nil

			for _, cs := range co2Strings {
				if co2BitCount == nil {
					co2BitCount = make([]int, len(cs))
				}

				for j := 0; j < len(cs); j++ {
					switch cs[j : j+1] {
					case "0":
						co2BitCount[j]--
					case "1":
						co2BitCount[j]++
					default:
						panic("something's wrong")
					}
				}
			}
		}

		co2 := []string{}

		switch {
		case co2BitCount[i] > 0:
			for k := range co2Strings {
				if co2Strings[k][i:i+1] == "0" {
					if len(co2Strings) > 1 {
						co2 = append(co2, co2Strings[k])
					}
				}
			}
		case co2BitCount[i] < 0:
			for k := range co2Strings {
				if co2Strings[k][i:i+1] == "1" {
					if len(co2Strings) > 1 {
						co2 = append(co2, co2Strings[k])
					}
				}
			}
		case co2BitCount[i] == 0:
			for k := range co2Strings {
				if co2Strings[k][i:i+1] == "0" {
					if len(co2Strings) > 1 {
						co2 = append(co2, co2Strings[k])
					}
				}
			}
		}

		log.Println("co2 is equal to...")
		log.Println(co2)

		if len(co2) > 0 {
			co2Strings = sliceutil.IntersectStrings(co2Strings, co2)
		}

		log.Println("co2Strings is now equal to...")
		log.Println(co2Strings)
		log.Println("")
	}

	oxDec, err := strconv.ParseInt(oxStrings[0], 2, 64)
	co2Dec, err := strconv.ParseInt(co2Strings[0], 2, 64)

	fmt.Printf("oxygen: %s (%d)\n", oxStrings[0], oxDec)
	fmt.Printf("co2: %s (%d)\n", co2Strings[0], co2Dec)
	fmt.Printf("multiplied: %d\n", oxDec*co2Dec)
}
