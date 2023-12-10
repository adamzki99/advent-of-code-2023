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

		if index < 0 {
			return "", -1
		}

		currentChar := s[index]

		if currentChar == "." {
			return "", -1
		}

		return currentChar, index
	}
}

// Only allowed start in a positive direction
func createPotentialMomentums() []Momentum {

	potentialStartingMomentums := []Momentum{}

	potentialStartingMomentums = append(potentialStartingMomentums, Momentum{horisontalDir: 0, verticalDir: 1})
	potentialStartingMomentums = append(potentialStartingMomentums, Momentum{horisontalDir: 1, verticalDir: 0})

	return potentialStartingMomentums
}

func DoesPointTouchAnyPointInSet(p, stride int, s []int) bool {

	//Maybe it already is in the set
	if slices.Contains(s, p) {
		return true
	}

	reachingPoints := []int{p - stride, p + 1, p + stride, p - 1}

	for _, rP := range reachingPoints {

		if slices.Contains(s, rP) {
			return true
		}

	}

	return false

}

// Returns the number of loop crossings in a given directions
func FireBeam(longestLoop *[]int, groundTile, stride, puzzleMapLen int, m Momentum) int {

	//       -----------------
	// 		 |  o----------->|
	//       |               |
	//       |               |
	//  o--->|-------------->|
	//       |               |
	//       -----------------
	//              o-------------
	//
	// If a line from a point hits a uneven amount of loop-lines, it is in the loop.
	// If not, it is outside the loop
	//

	nextGroundTile := groundTile + m.horisontalDir + (m.verticalDir * stride)

	if -1 < nextGroundTile && nextGroundTile%stride != 0 && nextGroundTile < puzzleMapLen {

		if slices.Contains(*longestLoop, nextGroundTile) {
			return 1 + FireBeam(longestLoop, nextGroundTile, stride, puzzleMapLen, m)
		} else {
			return FireBeam(longestLoop, nextGroundTile, stride, puzzleMapLen, m)
		}

	}

	return 0

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

	// Get all points not in main loop
	groundTiles := []int{}
	for i := range puzzleMap {

		if !slices.Contains(longestLoop, i) {
			groundTiles = append(groundTiles, i)
		}

	}

	pointsEnclosedByLoop := 0

	// Determine if the line is inside or outside the loop
	for _, groundTile := range groundTiles {

		hits := FireBeam(&longestLoop, groundTile, stride, len(puzzleMap), Momentum{horisontalDir: 1, verticalDir: 0})

		if hits%2 == 0 {
			continue
		}

		hits = FireBeam(&longestLoop, groundTile, stride, len(puzzleMap), Momentum{horisontalDir: -1, verticalDir: 0})

		if hits%2 == 0 {
			continue
		}

		hits = FireBeam(&longestLoop, groundTile, stride, len(puzzleMap), Momentum{horisontalDir: 0, verticalDir: 1})

		if hits%2 == 0 {
			continue
		}

		hits = FireBeam(&longestLoop, groundTile, stride, len(puzzleMap), Momentum{horisontalDir: 0, verticalDir: -1})

		if hits%2 == 0 {
			continue
		}

		pointsEnclosedByLoop++

	}

	return pointsEnclosedByLoop

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
