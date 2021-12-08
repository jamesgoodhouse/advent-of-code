package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputFile = "input.txt"
)

type (
	digit struct {
		segments segments
	}

	segments []string
)

var (
	digits = map[int]digit{
		0: {[]string{"a", "b", "c", "e", "f", "g"}},
		1: {[]string{"c", "f"}},
		2: {[]string{"a", "c", "d", "e", "g"}},
		3: {[]string{"a", "c", "d", "f", "g"}},
		4: {[]string{"b", "c", "d", "f"}},
		5: {[]string{"a", "b", "d", "f", "g"}},
		6: {[]string{"a", "b", "d", "e", "f", "g"}},
		7: {[]string{"a", "c", "f"}},
		8: {[]string{"a", "b", "c", "d", "e", "f", "g"}},
		9: {[]string{"a", "b", "c", "d", "f", "g"}},
	}

	counts = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		t := scanner.Text()

		segmentsRaw := strings.Split(t, " | ")

		// signalPatterns := strings.Fields(segmentsRaw[0])
		outputs := strings.Fields(segmentsRaw[1])

		for _, output := range outputs {
			switch {
			case len(output) == digits[1].numSegments():
				counts[1]++
			case len(output) == digits[4].numSegments():
				counts[4]++
			case len(output) == digits[7].numSegments():
				counts[7]++
			case len(output) == digits[8].numSegments():
				counts[8]++
			}
		}
	}

	fmt.Printf("total num of 1s, 4s, 7s, and 8s: '%d'\n", counts[1]+counts[4]+counts[7]+counts[8])
}

func (d digit) numSegments() int     { return len(d.segments) }
func (s *segments) toString() string { return strings.Join(*s, "") }
