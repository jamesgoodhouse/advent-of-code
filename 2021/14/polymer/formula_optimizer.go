package polymer

type (
	PairInsertionRule struct {
		Pair          string
		InsertionChar string
	}

	FormulaOptimizer struct {
		PairInsertionRules map[string]*PairInsertionRule
		Template           string
	}
)

func NewFormulaOptimizer(template string, rules map[string]*PairInsertionRule) *FormulaOptimizer {
	return &FormulaOptimizer{
		PairInsertionRules: rules,
		Template:           template,
	}
}

func NewPairInsertionRule(pair, char string) *PairInsertionRule {
	return &PairInsertionRule{Pair: pair, InsertionChar: char}
}

func (pfo *FormulaOptimizer) RunPairInsertionProcess(numSteps int) *Polymer {
	template := pfo.Template
	counts := map[string]int{}
	stepCounts := map[string]int{}
	charCount := map[string]int{}

	for i := 0; i < len(template); i++ {
		charCount[string(template[i])]++

		if i < len(template)-1 {
			pair := template[i : i+2]
			counts[pair]++
			stepCounts[pair]++
		}
	}

	for step := 1; step <= numSteps; step++ {
		for pair, count := range counts {
			charToInsert := pfo.PairInsertionRules[pair].InsertionChar
			charCount[charToInsert] += count
			stepCounts[charToInsert+pair[1:]] += count
			stepCounts[pair[:1]+charToInsert] += count
			stepCounts[pair] -= count
			if stepCounts[pair] <= 0 {
				delete(stepCounts, pair)
			}
		}

		counts = make(map[string]int)
		for key, val := range stepCounts {
			counts[key] = val
		}
	}

	return NewPolymer(charCount)
}
