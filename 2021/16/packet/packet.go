package packet

import (
	"fmt"
	"strconv"
)

type (
	Packet struct {
		binary string
		hex    string
	}
)

func New(hex string) *Packet {
	p := &Packet{hex: hex}

	for _, hc := range hex {
		i, err := strconv.ParseUint(string(hc), 16, 64)
		if err != nil {
			panic(err)
		}
		p.binary += fmt.Sprintf("%04b", i)
	}

	return p
}

func (p *Packet) Version() (*uint64, error) {
	version, err := strconv.ParseUint(p.binary[:3], 2, 64)
	if err != nil {
		return nil, err
	}
	return &version, nil
}

func (p *Packet) TypeID() (*uint64, error) {
	typeID, err := strconv.ParseUint(p.binary[3:6], 2, 64)
	if err != nil {
		return nil, err
	}
	return &typeID, nil
}
