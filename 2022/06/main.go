package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jamesgoodhouse/advent-of-code/2022/helper"
)

const (
	messageMarkerSize = 14
	packetMarkerSize  = 4
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("no input file give")
		os.Exit(1)
	}

	scanner, file, err := helper.NewFileScanner(os.Args[1], bufio.ScanRunes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	charCount := 0
	foundStartOfMessage := false
	foundStartOfPacket := false
	messageMarker := ""
	packetMarker := ""
	startOfMessageCharCount := 0
	startOfPacketCharCount := 0

	for scanner.Scan() {
		if foundStartOfMessage && foundStartOfPacket {
			break
		}

		charCount += 1
		r := scanner.Text()

		if !foundStartOfPacket {
			packetMarker += r
		}

		if !foundStartOfPacket && len(packetMarker) == packetMarkerSize {
			// check for all unique chars
			if containsUnique(packetMarker) {
				foundStartOfPacket = true
				startOfPacketCharCount = charCount
				fmt.Printf("packet marker '%v' found at '%v'\n", packetMarker, startOfPacketCharCount)
				continue
			}

			// if not full of unique chars, lop that first one off and move along
			packetMarker = packetMarker[1:]
		}

		if !foundStartOfMessage {
			messageMarker += r
		}

		if !foundStartOfMessage && len(messageMarker) == messageMarkerSize {
			// check for all unique chars
			if containsUnique(messageMarker) {
				foundStartOfMessage = true
				startOfMessageCharCount = charCount
				fmt.Printf("message marker '%v' found at '%v'\n", messageMarker, startOfMessageCharCount)
				continue
			}

			// if not full of unique chars, lop that first one off and move along
			messageMarker = messageMarker[1:]
		}
	}
}

func containsUnique(s string) bool {
	m := make(map[rune]bool)
	for _, i := range s {
		if _, ok := m[i]; ok {
			return false
		}

		m[i] = true
	}

	return true
}
