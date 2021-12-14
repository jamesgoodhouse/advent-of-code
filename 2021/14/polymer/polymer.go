package polymer

import (
	"strings"
)

type (
	Polymer string
)

func New(s string) *Polymer {
	p := Polymer(s)
	return &p
}

func (p *Polymer) LeastCommonElement(charSet []string) (string, int) {
	var leastCommonChar string
	var leastCommonCount int

	for char, count := range p.CountCharacters(charSet) {
		if leastCommonCount == 0 || count < leastCommonCount {
			leastCommonChar = char
			leastCommonCount = count
		}
	}
	return leastCommonChar, leastCommonCount
}

func (p *Polymer) MostCommonElement(charSet []string) (string, int) {
	var mostCommonChar string
	var mostCommonCount int

	for char, count := range p.CountCharacters(charSet) {
		if count > mostCommonCount {
			mostCommonChar = char
			mostCommonCount = count
		}
	}
	return mostCommonChar, mostCommonCount
}

func (p *Polymer) CountCharacters(charSet []string) map[string]int {
	count := map[string]int{}
	for _, char := range charSet {
		count[char] = strings.Count(string(*p), char)
	}
	return count
}
