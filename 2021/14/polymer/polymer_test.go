package polymer_test

import (
	"testing"

	"github.com/jamesgoodhouse/advent-of-code/2021/14/polymer"
	"github.com/stretchr/testify/assert"
)

func TestRunPairInsertionProcess(t *testing.T) {
	rules := map[string]*polymer.PairInsertionRule{
		"BB": {"BB", "N"},
		"BC": {"BC", "B"},
		"BH": {"BH", "H"},
		"BN": {"BN", "B"},
		"CB": {"CB", "H"},
		"CC": {"CC", "N"},
		"CH": {"CH", "B"},
		"CN": {"CN", "C"},
		"HB": {"HB", "C"},
		"HC": {"HC", "B"},
		"HH": {"HH", "N"},
		"HN": {"HN", "C"},
		"NB": {"NB", "B"},
		"NC": {"NC", "B"},
		"NH": {"NH", "C"},
		"NN": {"NN", "C"},
	}
	charSet := []string{"B", "C", "H", "N"}
	template := "NNCB"
	pfo := polymer.NewFormulaOptimizer(template, rules)

	testCases := map[string]struct {
		count        int
		numSteps     int
		checkPolymer bool
		polymer      *polymer.Polymer
	}{
		"single_step": {
			count:        1,
			numSteps:     1,
			checkPolymer: true,
			polymer:      polymer.New("NCNBCHB"),
		},
		"two_steps": {
			count:        5,
			numSteps:     2,
			checkPolymer: true,
			polymer:      polymer.New("NBCCNBBBCBHCB"),
		},
		"three_steps": {
			count:        7,
			numSteps:     3,
			checkPolymer: true,
			polymer:      polymer.New("NBBBCNCCNBBNBNBBCHBHHBCHB"),
		},
		"four_steps": {
			count:        18,
			numSteps:     4,
			checkPolymer: true,
			polymer:      polymer.New("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"),
		},
		"ten_steps": {
			count:    1588,
			numSteps: 10,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			polymer := pfo.RunPairInsertionProcess(tc.numSteps)
			_, lc := polymer.LeastCommonElement(charSet)
			_, mc := polymer.MostCommonElement(charSet)
			assert.Equal(t, tc.count, mc-lc)
			if tc.checkPolymer {
				assert.Equal(t, tc.polymer, polymer)
			}
		})
	}
}
