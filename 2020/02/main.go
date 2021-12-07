package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
	pattern   = `(?P<min>\d+)-(?P<max>\d+) (?P<char>[a-z]+): (?P<password>[a-z]+)`
)

var (
	validPasswordCount int
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		re := regexp.MustCompile(pattern)
		groups := re.SubexpNames()
		matches := re.FindAllStringSubmatch(text, -1)[0]

		var (
			min, max       int
			char, password string
		)

		for gi := 1; gi < len(groups); gi++ {
			val := matches[gi]

			switch groups[gi] {
			case "min":
				if min, err = strconv.Atoi(val); err != nil {
					panic("not an int")
				}
			case "max":
				if max, err = strconv.Atoi(val); err != nil {
					panic("not an int")
				}
			case "char":
				char = val
			case "password":
				password = val
			}
		}

		count := strings.Count(password, char)
		if count >= min && count <= max {
			validPasswordCount++
		}
	}

	fmt.Println(validPasswordCount)
}
