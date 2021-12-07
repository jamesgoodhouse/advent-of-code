package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

var (
	dynamicFuel               bool
	fuelTotal, minPos, maxPos int
	positions                 []int
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbersRaw := strings.Split(scanner.Text(), ",")
		positions = make([]int, len(numbersRaw))

		for i := range numbersRaw {
			number, err := strconv.Atoi(numbersRaw[i])
			if err != nil {
				panic(err)
			}

			positions[i] = number
			if i == 0 {
				minPos = number
				maxPos = number
			} else {
				if number < minPos {
					minPos = number
				}
				if number > maxPos {
					maxPos = number
				}
			}
		}
	}

	dynamicFuel = true

	for i := minPos; i <= maxPos; i++ {
		fuelTotalForSpecificStep := 0

		for j := range positions {
			fuelCost := 1
			fuelAggr := 0
			posChange := int(math.Abs(float64(positions[j] - i)))

			for chg := 0; chg < posChange; chg++ {
				fuelAggr += fuelCost
				if dynamicFuel {
					fuelCost++
				}
			}

			fuelTotalForSpecificStep += fuelAggr
		}

		if i == 0 {
			fuelTotal = fuelTotalForSpecificStep
		} else if fuelTotalForSpecificStep < fuelTotal {
			fuelTotal = fuelTotalForSpecificStep
		}
	}

	fmt.Printf("fuel cost: %d\n", fuelTotal)
}
