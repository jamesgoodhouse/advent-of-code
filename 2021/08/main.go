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

	// 8       5     2     3     7   9      6      4    0      1  | 5     3     5     3
	// acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf

	//   0:      1:      2:      3:      4:
	//  aaaa    ....    aaaa    aaaa    ....
	// b    c  .    c  .    c  .    c  b    c
	// b    c  .    c  .    c  .    c  b    c
	//  ....    ....    dddd    dddd    dddd
	// e    f  .    f  e    .  .    f  .    f
	// e    f  .    f  e    .  .    f  .    f
	//  gggg    ....    gggg    gggg    ....

	//   5:      6:      7:      8:      9:
	//  aaaa    aaaa    aaaa    aaaa    aaaa
	// b    .  b    .  .    c  b    c  b    c
	// b    .  b    .  .    c  b    c  b    c
	//  dddd    dddd    ....    dddd    dddd
	// .    f  e    f  .    f  e    f  .    f
	// .    f  e    f  .    f  e    f  .    f
	//  gggg    gggg    ....    gggg    gggg

	//   0:      1:      2:      3:      4:
	//  ....    ....    ....    ....    ....
	// .    .  .    a  .    .  .    .  e    a
	// .    .  .    a  .    .  .    .  e    a
	//  ....    ....    ....    ....    ffff
	// .    .  .    b  .    .  .    .  .    b
	// .    .  .    b  .    .  .    .  .    b
	//  ....    ....    ....    ....    ....

	//   5:      6:      7:      8:      9:
	//  ....    ....    dddd    dddd    ....
	// .    .  .    .  .    a  e    a  .    .
	// .    .  .    .  .    a  e    a  .    .
	//  ....    ....    ....    ffff    ....
	// .    .  .    .  .    b  g    b  .    .
	// .    .  .    .  .    b  g    b  .    .
	//  ....    ....    ....    cccc    ....

	//  0
	// 1 2
	//  3
	// 4 5
	//  6

	//  1
	//  .
	// . *
	//  .
	// . *
	//  .

	//  2
	//  *
	// . *
	//  *
	// * .
	//  *

	//  dddd
	// e    a
	// e    a
	//  ffff
	// g    b
	// g    b
	//  cccc

	// 8       3     3      4
	// acedgfa fbcad fbcad  eafb

	// 8       3     3      4
	// fdgacbe cefdb cefbgd gcbe

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
