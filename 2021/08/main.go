package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	inputFile = "input3.txt"
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

var numberSegmentCounts = map[int]int{
	1: 2,
	4: 4,
	7: 3,
	8: 7,
}

var numberSegments = map[int][]bool{
	0: {true, true, true, false, true, true, true},
	1: {false, false, true, false, false, true, false},
	2: {true, false, true, true, true, false, true},
	3: {true, false, true, true, false, true, true},
	4: {false, true, true, true, false, true, false},
	5: {true, true, false, true, false, true, true},
	6: {true, true, false, true, true, true, true},
	7: {true, false, true, false, false, true, false},
	8: {true, true, true, true, true, true, true},
	9: {true, true, true, true, false, true, true},
}

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

		knownNumberSegments := map[int]string{}

		for _, output := range outputs {
			// output := strings.Split(outputString, "")
			switch {
			case len(output) == numberSegmentCounts[1]:
				knownNumberSegments[1] = output
			case len(output) == numberSegmentCounts[4]:
				knownNumberSegments[4] = output
			case len(output) == numberSegmentCounts[7]:
				knownNumberSegments[7] = output
			case len(output) == numberSegmentCounts[8]:
				knownNumberSegments[8] = output
			}
		}

		permsOf8 := map[string]map[string]int{}
		Perm([]rune(knownNumberSegments[8]), func(a []rune) {
			permsOf8[string(a)] = make(map[string]int)
			for i, v := range strings.Split(string(a), "") {
				permsOf8[string(a)][v] = i
			}
		})

		for k, v := range knownNumberSegments {
			if k != 8 {
				var found bool
				fmt.Printf("checking number '%d' with segements '%s'\n", k, v)
				for _, p8 := range permsOf8 {
					knownNumber := make([]bool, 7)
					for _, taco := range strings.Split(v, "") {
						knownNumber[p8[taco]] = true
					}

					if reflect.DeepEqual(numberSegments[k], knownNumber) {
						fmt.Printf("permutation of '8' that fits '%d': %v\n", k, p8)
						fmt.Println(knownNumber)
						fmt.Println(numberSegments[k])
						found = true
						// break
					}
				}
				if !found {
					fmt.Printf("no fit found for '%d' :-(\n", k)
				}
				fmt.Println("--------------------")
			}
		}
	}
}

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
