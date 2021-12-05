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
		numbers                  [][]string
		matches                  [][]bool
		rowCounters, colCounters []int
	}
)

func newBoard() *board {
	b := &board{}
	b.numbers = make([][]string, 5)
	b.matches = make([][]bool, 5)
	for i := range b.numbers {
		b.numbers[i] = make([]string, 5)
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

	var numbers []string
	inputLine := 0

	boardLines := 0

	boards := []*board{}

	var b *board

	for scanner.Scan() {
		text := scanner.Text()

		if inputLine == 0 {
			numbers = strings.Split(text, ",")
		} else if text != "" {
			if b == nil {
				b = newBoard()
			}
			b.numbers[boardLines] = strings.Fields(text)
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
				n, err := strconv.Atoi(b.numbers[i][j])
				if err != nil {
					panic("something broke")
				}
				sum += n
			}
		}
	}
	return sum
}

func runBingo(numbers []string, boards []*board) (*board, *int) {
	for _, n := range numbers {
		for _, b := range boards {
			for bRowNum := 0; bRowNum < 5; bRowNum++ {
				for bColNum := 0; bColNum < 5; bColNum++ {
					if b.numbers[bRowNum][bColNum] == n {
						b.matches[bRowNum][bColNum] = true
						b.rowCounters[bRowNum]++
						b.colCounters[bColNum]++
						if b.rowCounters[bRowNum] == 5 || b.colCounters[bColNum] == 5 {
							winNum, err := strconv.Atoi(n)
							if err != nil {
								panic("something broke")
							}
							return b, &winNum
						}
					}
				}
			}
		}
	}

	return nil, nil
}
