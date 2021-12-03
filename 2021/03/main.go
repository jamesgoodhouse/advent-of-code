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
	inputFile = "input2.txt"
)

func main() {
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

	oxStrings := bitStrings

	var gamma string
	var epsilon string

	for i := range bitCount {
		log.Println("oxStrings is equal to...")
		log.Println(oxStrings)

		ox := []string{}

		switch {
		case bitCount[i] < 0:
			log.Printf("bit in place '%d' is more commonly 0\n", i)
			gamma = gamma + "0"
			epsilon = epsilon + "1"

			for j := range oxStrings {
				if oxStrings[j][i:i+1] == "0" {
					if len(oxStrings) > 1 {
						ox = append(ox, oxStrings[j])
					}
				}
			}
		case bitCount[i] > 0:
			log.Printf("bit in place '%d' is more commonly 1\n", i)
			gamma = gamma + "1"
			epsilon = epsilon + "0"

			for j := range oxStrings {
				if oxStrings[j][i:i+1] == "1" {
					if len(oxStrings) > 1 {
						ox = append(ox, oxStrings[j])
					}
				}
			}
		case bitCount[i] == 0:
			log.Printf("bit in place '%d' is equally 0 and 1", i)
			for j := range oxStrings {
				if oxStrings[j][i:i+1] == "1" {
					if len(oxStrings) > 1 {
						ox = append(ox, oxStrings[j])
					}
				}
			}
		default:
			panic("something's wrong")
		}

		if i == 4 {
			log.Println(bitCount)
		}

		log.Println("ox is equal to...")
		log.Println(ox)

		if len(ox) > 0 {
			// oxStrings = Stringify(intersect.Simple(oxStrings, ox))
			oxStrings = sliceutil.IntersectStrings(oxStrings, ox)
		}

		log.Println("oxStrings is now equal to...")
		log.Println(oxStrings)
		log.Println("")
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, err := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("gamma: %s (%d)\n", gamma, gammaDec)
	fmt.Printf("epsilon: %s (%d)\n", epsilon, epsilonDec)
	fmt.Printf("multiplied together: %d\n", epsilonDec*gammaDec)

	fmt.Println(oxStrings)
	// fmt.Printf("oxygen: %s (%d)\n", oxStrings, epsilonDec)
}
