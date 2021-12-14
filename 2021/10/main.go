package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

const (
	inputFile         = "input-test-2.txt"
	closeCharsPattern = `[}>)\]]{1}`
	openCharsPattern  = `[{<([]{1}`
)

var (
	closeCharReg = regexp.MustCompile(closeCharsPattern)
	openCharReg  = regexp.MustCompile(openCharsPattern)

	corruptedLines = []string{}

	incompleteLines = []string{}

	scores = []int{}

	matchingChars = map[string]string{
		"{": "}",
		"<": ">",
		"[": "]",
		"(": ")",

		"}": "{",
		">": "<",
		")": "(",
		"]": "[",
	}

	illegalCharCount = map[string]int{
		")": 0,
		"]": 0,
		"}": 0,
		">": 0,
	}

	illegalCharPoints = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
)

type (
	NavigationSubsystem struct {
		Lines []string
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
		CheckLine(scanner.Text())
	}

	fmt.Println(corruptedLines)
	fmt.Println(illegalCharCount)
	fmt.Println(
		illegalCharCount["}"]*illegalCharPoints["}"] + illegalCharCount[">"]*illegalCharPoints[">"] + illegalCharCount[")"]*illegalCharPoints[")"] + illegalCharCount["]"]*illegalCharPoints["]"],
	)
	fmt.Println("-----------------------")
	fmt.Println(incompleteLines)

	for _, l := range incompleteLines {
		ProcessIncompleteLine(l)
	}
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(len(scores) / 2)
	fmt.Println(scores[len(scores)/2])
}

func ProcessIncompleteLine(line string) {
	lineStack := strings.Split(line, "")
	charCount := map[string]int{}
	fmt.Println("PROCESING LINE")
	lineEndStack := []string{}
	fmt.Println(line)
	for _, char := range lineStack {
		charCount[char]++
	}

	for _, char := range lineStack {
		// fmt.Println(char)
		// fmt.Println(charCount[matchingChars[char]])
		// fmt.Println(charCount[char])
		if openCharReg.MatchString(char) &&
			charCount[matchingChars[char]] < charCount[char] {
			fmt.Println(char)
			lineEndStack = append([]string{matchingChars[char]}, lineEndStack...)
			charCount[matchingChars[char]]++
		}
	}
	fmt.Println(lineEndStack)

	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	score := 0
	scoreMultiplier := 5

	for _, endChar := range lineEndStack {
		score = score * scoreMultiplier
		score = score + points[endChar]
	}
	fmt.Printf("score %d\n", score)
	scores = append(scores, score)
}

func CheckLine(line string) {
	stack := []string{}
	corruptLine := false

	for _, currentChar := range strings.Split(line, "") {
		fmt.Println(stack)
		stack = append(stack, currentChar)
		// fmt.Println(stack)

		if len(stack) > 1 {
			if closeCharReg.MatchString(currentChar) {
				previousChar := (stack[len(stack)-2 : len(stack)-1])[0]
				fmt.Printf("found closing char '%s'\n", currentChar)
				fmt.Printf("previous char: '%s'\n", previousChar)
				fmt.Printf("expected previous char: '%s'\n", matchingChars[previousChar])
				if previousChar != matchingChars[currentChar] {
					corruptedLines = append(corruptedLines, line)
					fmt.Printf("corrupt line: '%s'\n", line)
					illegalCharCount[currentChar]++
					corruptLine = true
					break
				} else {
					fmt.Println("removing matching set from stack")
					stack = stack[:len(stack)-2]
					fmt.Println(stack)
				}
			}
		}
		fmt.Println("----------")
	}

	if !corruptLine {
		incompleteLines = append(incompleteLines, line)
	}
}

// [({(<(())[]>[[{[]{<()<>>
// [(()[<>])]({[<{<<[]>>(
// {([(<{}[<>[]}>{[]{[(<()>
// (((({<>}<{<{<>}{[]{[]{}
// [[<[([]))<([[{}[[()]]]
// [{[{({}]{}}([{[{{{}}([]
// {<[[]]>}<{[{[{[]{()[[[]
// [<(<(<(<{}))><([]([]()
// <{([([[(<>()){}]>(<<{{
// <{([{{}}[<[[[<>{}]]]>[]]
