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
	numSteps  = 40
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

	lineNum := 1
	for scanner.Scan() {
		if lineNum == 1 {
			template = scanner.Text()
		} else if lineNum > 2 {
			rawRule := strings.Split(scanner.Text(), " -> ")
			rules[rawRule[0]] = polymer.NewPairInsertionRule(rawRule[0], rawRule[1])
		}
		lineNum++
	}

	pfo := polymer.NewFormulaOptimizer(template, rules)
	polymer := pfo.RunPairInsertionProcess(numSteps)
	_, lc := polymer.LeastCommonElement()
	_, mc := polymer.MostCommonElement()

	fmt.Println(mc - lc)
}
