package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

type (
	board struct {
		numbers                  [][]int
		matches                  [][]bool
		rowCounters, colCounters []int
	}
)

func newBoard() *board {
	b := &board{}
	b.numbers = make([][]int, 5)
	b.matches = make([][]bool, 5)
	for i := range b.numbers {
		b.numbers[i] = make([]int, 5)
		b.matches[i] = make([]bool, 5)
		b.rowCounters = make([]int, 5)
		b.colCounters = make([]int, 5)
	}
	return b
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := []int{}
	inputLine := 0

	boardLines := 0

	boards := []*board{}

	var b *board

	for scanner.Scan() {
		text := scanner.Text()

		if inputLine == 0 {
			for _, numRaw := range strings.Split(text, ",") {
				n, err := strconv.Atoi(numRaw)
				if err != nil {
					panic("not a number")
				}
				numbers = append(numbers, n)
			}
		} else if text != "" {
			if b == nil {
				b = newBoard()
			}
			for i, numRaw := range strings.Fields(text) {
				n, err := strconv.Atoi(numRaw)
				if err != nil {
					panic("not a number")
				}
				b.numbers[boardLines][i] = n
			}
			boardLines++
		}

		if boardLines == 5 {
			boards = append(boards, b)
			b = nil
			boardLines = 0
		}

		inputLine++
	}

	winningBoard, winningNumber := runBingo(numbers, boards)
	if winningBoard == nil {
		panic("no winning board found")
	}

	sum := sumUnmarkedNums(winningBoard)

	fmt.Println(sum * *winningNumber)
}

func sumUnmarkedNums(b *board) int {
	sum := 0
	for i := range b.matches {
		for j := range b.matches[i] {
			if !b.matches[i][j] {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum
}

func runBingo(numbers []int, boards []*board) (*board, *int) {
	for _, n := range numbers {
		for _, b := range boards {
			for bRowNum := 0; bRowNum < 5; bRowNum++ {
				for bColNum := 0; bColNum < 5; bColNum++ {
					if b.numbers[bRowNum][bColNum] == n {
						b.matches[bRowNum][bColNum] = true
						b.rowCounters[bRowNum]++
						b.colCounters[bColNum]++
						if b.rowCounters[bRowNum] == 5 || b.colCounters[bColNum] == 5 {
							return b, &n
						}
					}
				}
			}
		}
	}

	return nil, nil
}
