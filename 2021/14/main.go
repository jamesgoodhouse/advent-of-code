package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jamesgoodhouse/advent-of-code/2021/14/polymer"
)

const (
	inputFile = "input.txt"
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
			charSetMap[char] = true
		}
	}

	charSet := []string{}
	for char := range charSetMap {
		charSet = append(charSet, char)
	}

	pfo := polymer.NewFormulaOptimizer(template, rules)
	polymer := pfo.RunPairInsertionProcess(10)
	_, lc := polymer.LeastCommonElement(charSet)
	_, mc := polymer.MostCommonElement(charSet)
	fmt.Println(mc - lc)
}
