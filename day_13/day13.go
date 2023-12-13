package main

import (
	"fmt"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

// Returns a slice of indexes, the values are the left column of the reflection
func FindColumnWithReflection(puzzle_input *[]string) []int {

	mirrorIndex := []int{}

columnLoop:
	for c := 0; c < len((*puzzle_input)[0])-1; c++ {

		for _, line := range *puzzle_input {

			if line[c] != line[c+1] {
				continue columnLoop
			}

		}

		mirrorIndex = append(mirrorIndex, c+1)
	}

	return mirrorIndex
}

// Returns a slice of indexes, the values are the top column of the reflection
func FindRowWithReflection(puzzle_input *[]string) []int {

	mirrorIndex := []int{}

rowLoop:
	for r := 0; r < len(*puzzle_input)-1; r++ {

		for c := 0; c < len((*puzzle_input)[r]); c++ {

			if (*puzzle_input)[r][c] != (*puzzle_input)[r+1][c] {
				continue rowLoop
			}

		}

		mirrorIndex = append(mirrorIndex, r+1)
	}

	return mirrorIndex
}

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
