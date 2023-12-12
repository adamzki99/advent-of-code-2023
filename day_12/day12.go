package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Arrangement struct {
	local        []string
	matchesUnder int
	operational  *Arrangement
	damaged      *Arrangement
}

func GetPattern(record []string) []int {

	patterns := []int{}

	brokenCounter := 0

	for _, c := range record {

		if c == "." && brokenCounter != 0 {
			patterns = append(patterns, brokenCounter)
			brokenCounter = 0
		}

		if c == "#" {
			brokenCounter++
		}

	}

	if brokenCounter != 0 {
		patterns = append(patterns, brokenCounter)
	}

	return patterns

}

func ExtractPattern(line string) []int {

	patternSlice := []int{}

	pattern := strings.Split(line, " ")[1]

	for _, c := range strings.Split(pattern, ",") {

		v, e := strconv.Atoi(c)

		if e == nil {
			patternSlice = append(patternSlice, v)
		}

	}

	return patternSlice
}

func PatternMatch(record []string, pattern []int) bool {

	recordPattern := GetPattern(record)

	if len(recordPattern) != len(pattern) {
		return false
	}

	for i := range recordPattern {

		if recordPattern[i] != pattern[i] {
			return false
		}

	}

	return true

}

func GenerateArrangements(a *Arrangement, pattern []int) {

	if !slices.Contains(a.local, "?") {

		if PatternMatch(a.local, pattern) {
			a.matchesUnder = 1
		} else {
			a.matchesUnder = 0
		}

		return
	}

	indexOfChange := slices.Index(a.local, "?")

	a.local[indexOfChange] = "#"
	localChange := []string{}
	localChange = append(localChange, a.local...)
	a.damaged = &Arrangement{local: localChange}
	GenerateArrangements(a.damaged, pattern)

	a.local[indexOfChange] = "."
	localChange = []string{}
	localChange = append(localChange, a.local...)
	a.operational = &Arrangement{local: localChange}
	GenerateArrangements(a.operational, pattern)

	// undo change
	a.local[indexOfChange] = "?"
	a.matchesUnder = a.damaged.matchesUnder + a.operational.matchesUnder
}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileLines := strings.Split(fileContent, "\n")

	puzzleAwnser := 0

	for _, line := range fileLines {

		pattern := ExtractPattern(line)

		lA := strings.Split(line, " ")[0]
		lineArrangement := strings.Split(lA, "")

		a := &Arrangement{local: lineArrangement}

		GenerateArrangements(a, pattern)

		puzzleAwnser = puzzleAwnser + a.matchesUnder

	}

	return puzzleAwnser

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
