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
	lanterfishTimerReset    = 6
	daysOfSpawning          = 80
)

var (
	fish = lanternfishes{}
)

type (
	lanternfish struct{ timer int }

	lanternfishes []*lanternfish
)

func newLanternfish(timer int) *lanternfish {
	return &lanternfish{timer: timer}
}

func (lf *lanternfish) createNewLaternfish() *lanternfish {
	return &lanternfish{timer: defaultLanternfishTimer}
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		for _, timerRaw := range strings.Split(text, ",") {
			timer, err := strconv.Atoi(timerRaw)
			if err != nil {
				panic(err)
			}
			fish = append(fish, &lanternfish{timer: timer})
		}
	}

	fish.spawn()

	fmt.Printf("number of lanternfish: '%d'\n", len(fish))
}

func (lf *lanternfishes) spawn() {
	for day := 0; day < daysOfSpawning; day++ {
		for _, f := range *lf {
			if f.timer == 0 {
				f.timer = 6
				fish = append(fish, &lanternfish{timer: defaultLanternfishTimer})
			} else if f.timer < 0 {
				panic("uhh...")
			} else {
				f.timer--
			}
		}

		fmt.Printf("day %d: %d\n", day, len(fish))
	}
}
