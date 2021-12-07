package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile               = "input.txt"
	defaultLanternfishTimer = 8
	lanternfishTimerReset   = 6
	daysOfSpawning          = 256
)

var (
	fish lanternfishes
)

type (
	lanternfishes []int
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		fish = make([]int, defaultLanternfishTimer+1)

		for _, timerRaw := range strings.Split(text, ",") {
			timer, err := strconv.Atoi(timerRaw)
			if err != nil {
				panic(err)
			}
			fish[timer]++
		}
	}

	fish.spawn()

	fmt.Printf("number of lanternfish: '%d'\n", fish.count())
}

func (lfs *lanternfishes) count() int {
	count := 0
	for _, lf := range *lfs {
		count += lf
	}
	return count
}

func (lf *lanternfishes) spawn() {
	for day := 0; day < daysOfSpawning; day++ {
		zeroFish := fish[0]
		fish = fish[1:]
		fish[lanternfishTimerReset] += zeroFish
		fish = append(fish, zeroFish)
	}
}
