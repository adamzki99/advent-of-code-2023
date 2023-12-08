package main

import (
	"fmt"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fmt.Println(fileContent)

	return -1

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
