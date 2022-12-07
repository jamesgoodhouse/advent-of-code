package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jamesgoodhouse/advent-of-code/2022/helper"
)

type ()

var (
	currentDir    = []string{}
	currentDirKey = ""

	dirSizes = map[string]int{}

	spaceNeededForUpdate = 30000000
	totalDiskSpace       = 70000000
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("no input file give")
		os.Exit(1)
	}

	scanner, file, err := helper.NewFileScanner(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "$ cd "): // is a `cd` command
			if !strings.HasPrefix(line, "$ cd ..") {
				currentDir = append(currentDir, strings.Split(line, "$ cd ")[1])
			} else {
				currentDir = currentDir[:len(currentDir)-1]
			}
			currentDirKey = fmt.Sprintf("%s%s", currentDir[0], strings.Join(currentDir[1:], "/"))
		case strings.HasPrefix(line, "$ ls"):
			// don't care about these currently
		case !strings.HasPrefix(line, "$") && !strings.HasPrefix(line, "dir"): // gross but it works
			sizeAndFile := strings.Split(line, " ")
			size, err := strconv.Atoi(sizeAndFile[0])
			if err != nil {
				panic("uhh....")
			}

			for i := range currentDir {
				dirSizes[fmt.Sprintf("%s%s", currentDir[0], strings.Join(currentDir[1:i+1], "/"))] += size
			}
		}
	}

	freeSpace := totalDiskSpace - dirSizes["/"]
	spaceToDelete := spaceNeededForUpdate - freeSpace

	sum := 0
	taco := 111111111111111111
	for _, size := range dirSizes {
		if size <= 100000 {
			sum += size
		}

		if size >= spaceToDelete && size < taco {
			taco = size
		}
	}

	fmt.Printf("sum of dirs <= 100000: %v\n", sum)
	fmt.Printf("free space: %v\n", freeSpace)
	fmt.Printf("space needed for update: %v\n", spaceToDelete)

	fmt.Println(taco)
}
