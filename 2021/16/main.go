package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jamesgoodhouse/advent-of-code/2021/16/converter"
)

const (
	inputFile = "input.txt"
	numSteps  = 40
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hexRaw string

	for scanner.Scan() {
		hexRaw = scanner.Text()
		break
	}

	c := converter.New()
	fmt.Println(c.ConvertHexToBinary(hexRaw))
}
