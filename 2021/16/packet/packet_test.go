package packet_test

import (
	"testing"

	"github.com/jamesgoodhouse/advent-of-code/2021/16/packet"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	testCases := map[string]struct {
		hex     string
		version uint64
	}{
		"something": {
			hex:     "D2FE28",
			version: 6,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			p := packet.New(tc.hex)
			version, _ := p.Version()
			assert.Equal(t, tc.version, *version)
		})
	}
}

func TestTypeID(t *testing.T) {
	testCases := map[string]struct {
		hex    string
		typeID uint64
	}{
		"something": {
			hex:    "D2FE28",
			typeID: 4,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			p := packet.New(tc.hex)
			typeID, _ := p.TypeID()
			assert.Equal(t, tc.typeID, *typeID)
		})
	}
}
