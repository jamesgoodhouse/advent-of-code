package polymer

import "fmt"

type (
	PairInsertionRule struct {
		Pair          string
		InsertionChar string
	}

	FormulaOptimizer struct {
		PairInsertionRules map[string]*PairInsertionRule
		Template           string
		Polymer            *Polymer
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
			fmt.Println(len(stepTemplate))
		}
		// fmt.Println(len(stepTemplate))
		template = stepTemplate
	}

	return New(template)
}
