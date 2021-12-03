package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

func main() {
	var aim, depth, hPos int

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
			aim = aim - value
		case "down":
			aim = aim + value
		case "forward":
			hPos = hPos + value
			depth = depth + aim*value
		default:
			panic("no idea what's going on")
		}
	}

	log.Printf("Horizontal Position: %d", hPos)
	log.Printf("Depth: %d", depth)
	log.Printf("Multiplied Together: %d", depth*hPos)
}
