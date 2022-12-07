package helper

import (
	"bufio"
	"os"
)

func NewFileScanner(filepath string, splitFn bufio.SplitFunc) (*bufio.Scanner, *os.File, error) {
	inputFile, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(splitFn)

	return fileScanner, inputFile, nil
}
