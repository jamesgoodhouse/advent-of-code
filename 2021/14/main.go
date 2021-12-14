package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputFile = "input.txt"
)

type (
	PairInsertionRule struct {
		Pair          string
		InsertionChar string
	}

	PolymerFormulaOptimizer struct {
		PairInsertionRules map[string]*PairInsertionRule
		Template           string
		Polymer            *Polymer
	}

	Polymer string
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var template string
	rules := map[string]*PairInsertionRule{}

	lineNum := 1
	for scanner.Scan() {
		if lineNum == 1 {
			template = scanner.Text()
		} else if lineNum > 2 {
			rawRule := strings.Split(scanner.Text(), " -> ")
			rules[rawRule[0]] = NewPairInsertionRule(rawRule[0], rawRule[1])
		}
		lineNum++
	}

	pfo := NewPolymerFormulaOptimizer(template, rules)
	polymer := pfo.RunPairInsertionProcess(10)
	_, lc := polymer.LeastCommonElement()
	_, mc := polymer.MostCommonElement()
	fmt.Println(mc - lc)
}

func NewPolymerFormulaOptimizer(template string, rules map[string]*PairInsertionRule) *PolymerFormulaOptimizer {
	return &PolymerFormulaOptimizer{
		PairInsertionRules: rules,
		Template:           template,
	}
}

func NewPairInsertionRule(pair, char string) *PairInsertionRule {
	return &PairInsertionRule{Pair: pair, InsertionChar: char}
}

func NewPolymer(s string) *Polymer {
	p := Polymer(s)
	return &p
}

func (p *Polymer) LeastCommonElement() (string, int) {
	var leastCommonChar string
	var leastCommonCount int

	for char, count := range p.CountCharacters() {
		if leastCommonCount == 0 || count < leastCommonCount {
			leastCommonChar = char
			leastCommonCount = count
		}
	}
	return leastCommonChar, leastCommonCount
}

func (p *Polymer) MostCommonElement() (string, int) {
	var mostCommonChar string
	var mostCommonCount int

	for char, count := range p.CountCharacters() {
		if count > mostCommonCount {
			mostCommonChar = char
			mostCommonCount = count
		}
	}
	return mostCommonChar, mostCommonCount
}

func (p *Polymer) CountCharacters() map[string]int {
	count := map[string]int{}
	for _, char := range []string{"B", "C", "H", "N"} {
		count[char] = strings.Count(string(*p), char)
	}
	return count
}

func (pfo *PolymerFormulaOptimizer) RunPairInsertionProcess(numSteps int) *Polymer {
	template := pfo.Template

	for step := 1; step <= numSteps; step++ {
		stepTemplate := ""
		for i := 0; i < len(template)-1; i++ {
			pair := template[i : i+2]
			charToInsert := pfo.PairInsertionRules[pair].InsertionChar
			if i == 0 {
				stepTemplate += pair[:1] + charToInsert + pair[1:]
			} else {
				stepTemplate += charToInsert + pair[1:]
			}
		}
		template = stepTemplate
	}

	return NewPolymer(template)
}
