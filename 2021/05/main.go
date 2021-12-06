package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	inputFile = "input.txt"
	regString = `(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)`
)

type (
	point       struct{ x, y int }
	lineSegment struct{ p1, p2 *point }
	grid        [][]int
)

var (
	lineSegments []*lineSegment
	gridImpl     grid = func() [][]int {
		g := make([][]int, 1000)
		for i := range g {
			g[i] = make([]int, 1000)
		}
		return g
	}()
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

		l := &lineSegment{
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

		if l.isHorizontal() {
			if l.p1.x > l.p2.x {
				_p := l.p1.x
				l.p1.x = l.p2.x
				l.p2.x = _p
			}

			for x := l.p1.x; x <= l.p2.x; x++ {
				gridImpl[l.p1.y][x]++
			}
		} else if l.isVertical() {
			if l.p1.y > l.p2.y {
				_p := l.p1.y
				l.p1.y = l.p2.y
				l.p2.y = _p
			}

			for y := l.p1.y; y <= l.p2.y; y++ {
				gridImpl[y][l.p1.x]++
			}
		} else {
			x := l.p1.x
			y := l.p1.y

			for {
				gridImpl[y][x]++

				if x == l.p2.x || y == l.p2.y {
					break
				}

				if x >= l.p2.x {
					x--
				} else {
					x++
				}

				if y >= l.p2.y {
					y--
				} else {
					y++
				}
			}
		}

		lineSegments = append(lineSegments, l)
	}

	numOverlaps := 0
	for i := range gridImpl {
		for j := range gridImpl[i] {
			if gridImpl[i][j] >= 2 {
				numOverlaps++
			}
		}
	}

	fmt.Println(numOverlaps)
}

func (l lineSegment) isHorizontal() (b bool) {
	if l.p1.y == l.p2.y {
		b = true
	}
	return
}

func (l lineSegment) isVertical() (b bool) {
	if l.p1.x == l.p2.x {
		b = true
	}
	return
}
