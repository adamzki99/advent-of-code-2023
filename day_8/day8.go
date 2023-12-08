package main

import (
	"fmt"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Paths struct {
	left  string
	right string
}

func CharFactory(s string) func() string {
	index := 0
	return func() string {
		if index >= len(s) {
			index = 0
		}
		character := string(s[index])
		index++
		return character
	}
}

func LineFactory(s []string) func() (string, int) {
	index := 0
	return func() (string, int) {
		if index >= len(s) {
			index = 0
		}
		character := s[index]
		index++
		return character, (index - 1)
	}
}

func RebuildSliceWithIndexes(puzzleLines []string, toSave []int) []string {

	replacementSlice := []string{}

	for _, i := range toSave {

		replacementSlice = append(replacementSlice, puzzleLines[i])

	}

	return replacementSlice

}

func SplitNodeLine(puzzleLine string) (string, string, string) {

	puzzleLineSymbols := strings.Split(puzzleLine, "")

	return1 := puzzleLineSymbols[0] + puzzleLineSymbols[1] + puzzleLineSymbols[2]
	return2 := puzzleLineSymbols[7] + puzzleLineSymbols[8] + puzzleLineSymbols[9]
	return3 := puzzleLineSymbols[12] + puzzleLineSymbols[13] + puzzleLineSymbols[14]

	return return1, return2, return3
}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	places := make(map[string]Paths)

	puzzleLines := strings.Split(fileContent, "\n")

	nodeLines := puzzleLines[2:]

	for _, line := range nodeLines {

		p1, p2, p3 := SplitNodeLine(line)

		places[p1] = Paths{left: p2, right: p3}

	}

	stepFactory := CharFactory(puzzleLines[0])
	steps := 0

	currentPlace := "AAA"

	for {

		step := stepFactory()
		steps++

		if step == "L" {
			currentPlace = places[currentPlace].left
		} else {
			currentPlace = places[currentPlace].right
		}

		if currentPlace == "ZZZ" {
			fmt.Println(places[currentPlace])
			break
		}

	}

	return steps

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
