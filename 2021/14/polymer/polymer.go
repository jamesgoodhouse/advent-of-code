package polymer

type (
	Polymer struct {
		CharacterCounts map[string]int
	}
)

func NewPolymer(c map[string]int) *Polymer {
	return &Polymer{CharacterCounts: c}
}

func (p *Polymer) LeastCommonElement() (string, int) {
	var leastCommonChar string
	var leastCommonCount int

	for char, count := range (*p).CharacterCounts {
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

	for char, count := range (*p).CharacterCounts {
		if count > mostCommonCount {
			mostCommonCount = count
			mostCommonChar = char
		}
	}

	return mostCommonChar, mostCommonCount
}
