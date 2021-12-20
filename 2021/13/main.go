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

	dotPattern           = `(?P<x>\d+),(?P<y>\d+)`
	foldDirectionPattern = `fold along (?P<foldDirection>x|y|)=(?P<foldLine>\d+)`
)

type (
	Fold struct {
		Direction string
		Line      int
	}
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dots := make([][]bool, 1400)
	for i := 0; i < 1400; i++ {
		dots[i] = make([]bool, 1400)
	}

	dotReg := regexp.MustCompile(dotPattern)
	foldReg := regexp.MustCompile(foldDirectionPattern)

	folds := []Fold{}

	var maxX, maxY int

	for scanner.Scan() {
		if matches := dotReg.FindStringSubmatch(scanner.Text()); len(matches) > 0 {
			x, err := strconv.Atoi(matches[dotReg.SubexpIndex("x")])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(matches[dotReg.SubexpIndex("y")])
			if err != nil {
				panic(err)
			}
			if y > maxY {
				maxY = y
			}
			if x > maxX {
				maxX = x
			}
			dots[y][x] = true
		} else if matches := foldReg.FindStringSubmatch(scanner.Text()); len(matches) > 0 {
			foldLine, err := strconv.Atoi(matches[foldReg.SubexpIndex("foldLine")])
			if err != nil {
				panic(err)
			}
			folds = append(folds, Fold{
				matches[foldReg.SubexpIndex("foldDirection")],
				foldLine,
			})
		}
	}

	// gross but works to get the slice the correct size
	dots = dots[:maxY+1]
	for i := 0; i < len(dots); i++ {
		dots[i] = dots[i][:maxX+1]
	}

	for _, f := range folds {
		dots = fold(dots, f.Direction, f.Line)
	}

	numVisibleDots := 0

	for y := range dots {
		for x := range dots[y] {
			if dots[y][x] {
				fmt.Printf("#")
				numVisibleDots++
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println(numVisibleDots)
}

func fold(dots [][]bool, direction string, foldLine int) [][]bool {
	if direction == "y" {
		fmt.Printf("fold line is %s=%d\n", direction, foldLine)
		for i := foldLine + 1; i < len(dots); i++ {
			distanceFromFoldLine := i - foldLine
			toFoldTo := foldLine - distanceFromFoldLine
			if toFoldTo >= 0 {
				fmt.Printf("  folding line %d to %d\n", i, toFoldTo)
				for x := range dots[i] {
					if dots[i][x] && !dots[toFoldTo][x] {
						dots[toFoldTo][x] = dots[i][x]
					}
				}
			} else {
				fmt.Printf("  ignoring folding line %d\n", i)
			}
		}

		dots = dots[:foldLine]
	} else if direction == "x" {
		fmt.Printf("fold column is %s=%d\n", direction, foldLine)
		for y := range dots {
			fmt.Printf("  fold line is %d\n", y)
			for x := foldLine + 1; x < len(dots[y]); x++ {
				distanceFromFoldLine := x - foldLine
				toFoldTo := foldLine - distanceFromFoldLine
				if toFoldTo >= 0 {
					fmt.Printf("    folding column %d to %d\n", x, toFoldTo)
					if dots[y][x] && !dots[y][toFoldTo] {
						dots[y][toFoldTo] = dots[y][x]
					}
				} else {
					fmt.Printf("    ignoring folding column %d\n", x)
				}
			}
			dots[y] = dots[y][:foldLine]
		}
	}

	return dots
}
