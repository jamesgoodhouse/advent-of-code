package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input2.txt"
)

func main() {
	var depth, hPos int

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// https://stackoverflow.com/a/16615559/12369692
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")

		value, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err)
		}

		switch command[0] {
		case "up":
			depth = depth - value
		case "down":
			depth = depth + value
		case "forward":
			hPos = hPos + value
		default:
			panic("no idea what's going on")
		}
	}

	log.Println("Part One")
	log.Printf("  Horizontal Position: %d", hPos)
	log.Printf("  Depth: %d", depth)
	log.Printf("  Multiplied Together: %d", depth*hPos)
}
