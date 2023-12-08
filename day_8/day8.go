package main

import (
	"fmt"
	"math"
	"strings"
	"sync"

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

func Walk(startPlace string, walkPattern string, places map[string]Paths, wg *sync.WaitGroup, resultChan chan int) {

	defer wg.Done()
	stepFactory := CharFactory(walkPattern)
	steps := 0

	currentPlace := startPlace

	for {

		step := stepFactory()
		steps++

		if step == "L" {
			currentPlace = places[currentPlace].left
		} else {
			currentPlace = places[currentPlace].right
		}

		if currentPlace[2] == byte('Z') {
			break
		}

	}

	resultChan <- steps

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func LowesCommonMultipler(v []int) int {

	result := 1
	for _, num := range v {
		result = lcm(result, num)
	}
	return result

}

func SolvePuzzle(fileName string) int {

	var wg sync.WaitGroup

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	places := make(map[string]Paths)

	puzzleLines := strings.Split(fileContent, "\n")

	nodeLines := puzzleLines[2:]

	startingPlaces := []string{}

	for _, line := range nodeLines {

		p1, p2, p3 := SplitNodeLine(line)

		if p1[2] == byte('A') {
			startingPlaces = append(startingPlaces, p1)
		}

		places[p1] = Paths{left: p2, right: p3}

	}

	resultChan := make(chan int, len(startingPlaces))
	for _, sp := range startingPlaces {

		wg.Add(1)
		go Walk(sp, puzzleLines[0], places, &wg, resultChan)

	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := []int{}
	for r := range resultChan {

		results = append(results, r)
	}

	return LowesCommonMultipler(results)

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
