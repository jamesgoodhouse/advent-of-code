package helper

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func LoadFileFromArgs(args []string) (*os.File, error) {
	if len(args[1:]) > 1 {
		return nil, errors.New("too many args given")
	} else if len(args[1:]) != 1 {
		return nil, errors.New("no input file given")
	}
	return os.Open(args[1])
}

func NewFileScanner(file *os.File, splitFunc bufio.SplitFunc) (*bufio.Scanner, error) {
	if file == nil {
		return nil, errors.New("file cannot be nil")
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(splitFunc)
	return fileScanner, nil
}
