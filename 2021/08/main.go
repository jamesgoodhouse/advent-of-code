package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

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
//  T
// T T
//  F
// T T
//  T
// TTTFTTT

//  1
//  F
// F T
//  F
// F T
//  F
// FFTFFTF

//  2
//  T
// F T
//  T
// T F
//  T
// TFTTFT

//  3
//  T
// F T
//  T
// F T
//  T
// TFTTFTT

//  4
//  F
// T T
//  T
// F T
//  F
// FTTTFTF

//  5
//  T
// T F
//  T
// F T
//  T
// TTFTFTT

//  6
//  T
// T F
//  T
// T T
//  T
// TTFTTTT

//  7
//  T
// F T
//  F
// F T
//  F
// TFTFFTF

var (
	numberSegmentShapes = [][]bool{
		{true, true, true, false, true, true, true},
		{false, false, true, false, false, true, false},
		{true, false, true, true, true, false, true},
		{true, false, true, true, false, true, true},
		{false, true, true, true, false, true, false},
		{true, true, false, true, false, true, true},
		{true, true, false, true, true, true, true},
		{true, false, true, false, false, true, false},
		{true, true, true, true, true, true, true},
		{true, true, true, true, false, true, true},
	}

	outputValues int

	verbose = true
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

		signalPatterns := strings.Fields(segmentsRaw[0])
		outputs := strings.Fields(segmentsRaw[1])

		segments := []string{}
		permsOf8 := map[string]map[string]int{}

		for _, sp := range signalPatterns {
			if len(sp) == 7 {
				Perm([]rune(sp), func(a []rune) {
					permsOf8[string(a)] = make(map[string]int)
					for i, v := range strings.Split(string(a), "") {
						permsOf8[string(a)][v] = i
					}
				})
			} else {
				segments = append(segments, sp)
			}
		}

		knownSegments := make([]string, 10)
		allFound := false
		permIteration := 0

		for k, p8 := range permsOf8 {
			if verbose {
				fmt.Printf("permutation %d — '%s'\n", permIteration, k)
			}
			knownSegments[8] = k
			segmentIterator := 0

			for {
				segment := segments[segmentIterator]

				if verbose {
					fmt.Printf("checking segement '%s'\n", segment)
				}

				found := false

				numberShape := make([]bool, 7)

				for _, s := range strings.Split(segment, "") {
					numberShape[p8[s]] = true
				}

				for n, ns := range numberSegmentShapes {
					if n != 8 && reflect.DeepEqual(ns, numberShape) && knownSegments[n] == "" {
						if verbose {
							fmt.Printf("match found for '%d'\n", n)
						}
						knownSegments[n] = segment
						found = true
						break
					}
				}

				if found {
					segmentIterator++
				} else {
					if verbose {
						fmt.Printf("no match found for %s, starting over\n", segment)
						fmt.Println("---------------------------")
					}
					segmentIterator = 0
					knownSegments = make([]string, 10)
					break
				}

				if segmentIterator == 9 {
					allFound = true
					break
				}
			}
			if allFound {
				if verbose {
					fmt.Println("found matches for all segments!")
				}
				break
			}
			permIteration++
		}

		if !allFound {
			fmt.Println("uh... nope")
			os.Exit(0)
		}

		knownSegmentsNormalized := map[string]int{}
		for i, ks := range knownSegments {
			sorted := strings.Split(ks, "")
			sort.Strings(sorted)
			knownSegmentsNormalized[strings.Join(sorted, "")] = i
		}

		var outputValue string
		for _, o := range outputs {
			sorted := strings.Split(o, "")
			sort.Strings(sorted)
			outputValue += strconv.Itoa(knownSegmentsNormalized[strings.Join(sorted, "")])
		}

		ov, err := strconv.Atoi(outputValue)
		if err != nil {
			panic(err)
		}

		outputValues += ov
	}

	fmt.Printf("sum: %d\n", outputValues)
}

// https://yourbasic.org/golang/generate-permutation-slice-string/

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
