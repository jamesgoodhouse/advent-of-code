package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

const (
	inputFile = "input2.txt"
	regString = `(?P<x1>\d+),(?P<y2>\d+) -> (?P<x2>\d+),(?P<y2>\d+)`
)

type (
	point  struct{ x, y int }
	vector struct{ p1, p2 *point }
)

var (
	vectors []*vector
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

		re := regexp.MustCompile(regString)
		groups := re.SubexpNames()
		matches := re.FindAllStringSubmatch(text, -1)[0]

		l := &vector{
			p1: &point{},
			p2: &point{},
		}

		for gi := 1; gi < len(groups); gi++ {
			val, err := strconv.Atoi(matches[gi])
			if err != nil {
				panic(err)
			}

			switch groups[gi] {
			case "x1":
				l.p1.x = val
			case "y1":
				l.p1.y = val
			case "x2":
				l.p2.x = val
			case "y2":
				l.p2.y = val
			}
		}

		vectors = append(vectors, l)
	}
}

func (l vector) isVertical() (b bool) {
	if l.p1.x == l.p2.x {
		b = true
	}
	return
}

func (l vector) isHorizontal() (b bool) {
	if l.p1.y == l.p2.y {
		b = true
	}
	return
}
