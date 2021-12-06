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
	boardNumber struct {
		number int
		called bool
	}

	board struct {
		numbers                  [][]*boardNumber
		rowCounters, colCounters []int
		won                      bool
		winningNumber            int
	}

	winningBoards struct {
		boards []*board
	}
)

func newBoard() *board {
	b := &board{}
	b.numbers = make([][]*boardNumber, 5)
	for i := range b.numbers {
		b.numbers[i] = make([]*boardNumber, 5)
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
				b.numbers[boardLines][i] = &boardNumber{number: n}
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

	winBoards := runBingo(numbers, boards)

	firstWinBoard := winBoards.boards[0]
	fmt.Println(sumUnmarkedNums(firstWinBoard) * firstWinBoard.winningNumber)

	lastWinBoard := winBoards.boards[len(winBoards.boards)-1]
	fmt.Println(sumUnmarkedNums(lastWinBoard) * lastWinBoard.winningNumber)

}

func sumUnmarkedNums(b *board) int {
	sum := 0
	for i := range b.numbers {
		for j := range b.numbers[i] {
			if !b.numbers[i][j].called {
				sum += b.numbers[i][j].number
			}
		}
	}
	return sum
}

func runBingo(numbers []int, boards []*board) *winningBoards {
	winBoards := &winningBoards{}

	for _, n := range numbers {
		for _, b := range boards {
			if !b.won {
				for bRowNum := 0; bRowNum < 5; bRowNum++ {
					for bColNum := 0; bColNum < 5; bColNum++ {
						if b.numbers[bRowNum][bColNum].number == n {
							b.numbers[bRowNum][bColNum].called = true
							b.rowCounters[bRowNum]++
							b.colCounters[bColNum]++
							if b.rowCounters[bRowNum] == 5 || b.colCounters[bColNum] == 5 {
								b.won = true
								b.winningNumber = n
								winBoards.boards = append(winBoards.boards, b)
							}
						}
					}
				}
			}
		}
	}

	return winBoards
}
