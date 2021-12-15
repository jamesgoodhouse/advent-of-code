package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/jamesgoodhouse/advent-of-code/2021/14/polymer"
)

const (
	inputFile = "input.txt"
	numSteps  = 10
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var template string
	rules := map[string]*polymer.PairInsertionRule{}
	charSetMap := map[string]bool{}

	lineNum := 1
	for scanner.Scan() {
		if lineNum == 1 {
			template = scanner.Text()
		} else if lineNum > 2 {
			rawRule := strings.Split(scanner.Text(), " -> ")
			rules[rawRule[0]] = polymer.NewPairInsertionRule(rawRule[0], rawRule[1])
		}
		lineNum++

		for _, char := range strings.Split(scanner.Text(), "") {
			if unicode.IsLetter([]rune(char)[0]) {
				charSetMap[char] = true
			}
		}
	}

	charSet := []string{}
	for char := range charSetMap {
		charSet = append(charSet, char)
	}

	pfo := polymer.NewFormulaOptimizer(template, rules)
	polymer := pfo.RunPairInsertionProcess(numSteps)
	_, lc := polymer.LeastCommonElement(charSet)
	_, mc := polymer.MostCommonElement(charSet)
	fmt.Println(mc - lc)
}
