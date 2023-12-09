package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

func ProduceDifferences(historyReading *[]int) []int {

	rS := []int{}

	for i := 0; i < len(*historyReading)-1; i++ {

		rS = append(rS, (*historyReading)[i+1]-(*historyReading)[i])
	}

	return rS
}

func StringLineToIntSlice(sS *string) []int {

	rS := []int{}

	elements := strings.Split(*sS, " ")

	for _, e := range elements {
		intElement, err := strconv.Atoi(e)

		if err == nil {
			rS = append(rS, intElement)

		}
	}

	return rS
}

func CalculateNextValue(l1, l2 *[]int) {

	(*l1) = append((*l1), (*l2)[len((*l2))-1]+(*l1)[len((*l1))-1])

}

func AllElementsAreZero(l1 *[]int) bool {

	for _, e := range *l1 {
		if e != 0 {
			return false
		}
	}

	return true

}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	historyReadings := strings.Split(fileContent, "\n")

	puzzleResult := 0

	for _, hR := range historyReadings {

		differenceSlice := StringLineToIntSlice(&hR)

		iSlice := [][]int{}
		iSlice = append(iSlice, differenceSlice)

		for {
			differenceSlice = ProduceDifferences(&differenceSlice)
			iSlice = append(iSlice, differenceSlice)

			if AllElementsAreZero(&differenceSlice) {
				break
			}
		}

		slices.Reverse(iSlice)
		for i := 0; i < len(iSlice)-1; i++ {
			CalculateNextValue(&(iSlice[i+1]), &iSlice[i])
		}

		slices.Reverse(iSlice)
		puzzleResult = puzzleResult + iSlice[0][len(iSlice[0])-1]

	}

	return puzzleResult

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
