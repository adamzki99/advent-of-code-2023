package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Momentum struct {
	horisontalDir int // -1 = leftwards, 0 = still, 1 = rightwards
	verticalDir   int // -1 = uppwards, 0 = still, 1 = downwards
}

func NextPipeIndex(currentPipe string, currentIndex, stride int, m *Momentum) int {

	switch currentPipe {
	case "|":
		return currentIndex + (stride * m.verticalDir)
	case "-":
		return currentIndex + m.horisontalDir
	case "L":

		if m.horisontalDir != 0 {
			m.verticalDir = m.horisontalDir
			m.horisontalDir = 0

			return currentIndex + (stride * m.verticalDir)
		} else {
			m.horisontalDir = m.verticalDir
			m.verticalDir = 0

			return currentIndex + m.horisontalDir
		}

	case "J":

		if m.horisontalDir != 0 {
			m.verticalDir = m.horisontalDir * -1
			m.horisontalDir = 0

			return currentIndex + (stride * m.verticalDir)
		} else {
			m.horisontalDir = m.verticalDir * -1
			m.verticalDir = 0

			return currentIndex + m.horisontalDir
		}
	case "7":

		if m.horisontalDir != 0 {
			m.verticalDir = m.horisontalDir
			m.horisontalDir = 0

			return currentIndex + (stride * m.verticalDir)
		} else {
			m.horisontalDir = m.verticalDir
			m.verticalDir = 0

			return currentIndex + m.horisontalDir
		}
	case "F":

		if m.horisontalDir != 0 {
			m.verticalDir = m.horisontalDir * -1
			m.horisontalDir = 0

			return currentIndex + (stride * m.verticalDir)
		} else {
			m.horisontalDir = m.verticalDir * -1
			m.verticalDir = 0

			return currentIndex + m.horisontalDir
		}
	default:
		return -1
	}

}

func StepFactory(s []string, start, stride int, m *Momentum) func() (string, int) {
	index := start
	return func() (string, int) {

		index = NextPipeIndex(s[index], index, stride, m)

		if index == -1 {
			return "", -1
		}

		currentChar := s[index]

		if currentChar == "." {
			return "", -1
		}

		return currentChar, index
	}
}

func createPotentialMomentums() []Momentum {

	values := []int{1, -1}

	potentialStartingMomentums := []Momentum{}

	for _, v := range values {
		potentialStartingMomentums = append(potentialStartingMomentums, Momentum{horisontalDir: 0, verticalDir: v})

		potentialStartingMomentums = append(potentialStartingMomentums, Momentum{horisontalDir: v, verticalDir: 0})

	}

	return potentialStartingMomentums
}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	puzzleLines := strings.Split(fileContent, "\n")
	stride := len(puzzleLines[0])

	puzzleMap := []string{}
	for _, line := range puzzleLines {

		puzzleMap = append(puzzleMap, strings.Split(line, "")...)

	}

	// Find start point
	startIndex := slices.Index(puzzleMap, "S")

	if startIndex == -1 {
		fmt.Println("Could not find start index of map.")
		return -1
	}

	//Create starting environments
	potentialStartingPositions := []string{"|", "-"}
	potentialStartingMomentums := createPotentialMomentums()

	longestLoop := []int{}
	for _, pSS := range potentialStartingPositions {

		for _, pSM := range potentialStartingMomentums {

			startIndex := NextPipeIndex(pSS, startIndex, stride, &pSM)

			if startIndex == -1 || puzzleMap[startIndex] == "." { // Invalid start
				continue
			}

			stepFactory := StepFactory(puzzleMap, startIndex, stride, &pSM)

			loopLength := int16(0)

			currentPath := []int{startIndex}

			for {

				loopLength++

				nextPipe, pipeLocation := stepFactory()

				currentPath = append(currentPath, pipeLocation)

				if nextPipe == "" || nextPipe == "." { //Out of bounds, not a valid loop
					loopLength = -1
					break
				}

				if loopLength == math.MaxInt16 { // Probably ended up in a endless loop
					loopLength = -1
					break
				}

				if nextPipe == "S" {
					loopLength++
					break
				}

			}

			if len(longestLoop) < len(currentPath) && loopLength != -1 {
				longestLoop = currentPath
			}

		}

	}

	// Furthest point away is half of the loop lenght
	return len(longestLoop) / 2

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
