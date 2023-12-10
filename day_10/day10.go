package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"sync"

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

func StepFactory(s []string, start, stride int, m *Momentum) func() string {
	index := start
	return func() string {

		index = NextPipeIndex(s[index], index, stride, m)

		if index == -1 {
			return ""
		}

		currentChar := s[index]

		if currentChar == "." {
			return ""
		}

		return currentChar
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

func Walker(puzzleMap *[]string, startIndex, stride, id int, m *Momentum, wg *sync.WaitGroup, resultChan chan int) {

	defer wg.Done()

	stepFactory := StepFactory(*puzzleMap, startIndex, stride, m)

	loopLength := int16(0)

	for {

		loopLength++

		nextPipe := stepFactory()

		if nextPipe == "" || nextPipe == "." { //Out of bounds, not a valid loop
			loopLength = -1
			break
		}

		if loopLength == math.MaxInt16 {
			fmt.Printf("WARNING: Walker %d will hit overflow when counting loop lenght\n", id)
			loopLength = -1
			break
		}

		if nextPipe == "S" {
			loopLength++
			break
		}

	}
	resultChan <- int(loopLength)
}

func SolvePuzzle(fileName string) int {

	var wg sync.WaitGroup

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

	// Create a thread for each starting environment
	resultChan := make(chan int, len(potentialStartingPositions)*len(potentialStartingMomentums))
	workerID := -1
	for _, pSS := range potentialStartingPositions {

		for _, pSM := range potentialStartingMomentums {

			workerID++

			individualMomentum := Momentum{horisontalDir: pSM.horisontalDir, verticalDir: pSM.verticalDir}

			individualStartIndex := NextPipeIndex(pSS, startIndex, stride, &individualMomentum)

			if individualStartIndex == -1 || puzzleMap[individualStartIndex] == "." { // Invalid start
				continue
			}

			wg.Add(1)
			go Walker(&puzzleMap, individualStartIndex, stride, workerID, &individualMomentum, &wg, resultChan)

		}

	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Now get the longest loop
	results := []int{}
	for r := range resultChan {

		results = append(results, r)
	}

	slices.Sort(results)

	// Furthest point away is half of the loop lenght
	return results[len(results)-1] / 2

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
