package polymer

type (
	Polymer map[string]int
)

func New(s map[string]int) *Polymer {
	p := Polymer(s)
	return &p
}

func (p *Polymer) LeastCommonElement() (string, int) {
	var leastCommonChar string
	var leastCommonCount int

	for char, count := range *p {
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

	for char, count := range *p {
		if count > mostCommonCount {
			mostCommonCount = count
			mostCommonChar = char
		}
	}

	return mostCommonChar, mostCommonCount
}
