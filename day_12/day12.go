package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Arrangement struct {
	local []string
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

func ExtractPattern(line string, copies int) []int {

	patternSlice := []int{}

	pattern := strings.Split(line, " ")[1]

	for _, c := range strings.Split(pattern, ",") {

		v, e := strconv.Atoi(c)

		if e == nil {
			patternSlice = append(patternSlice, v)
		}

	}

	ogPatternSlice := []int{}
	ogPatternSlice = append(ogPatternSlice, patternSlice...)

	for i := 0; i < copies; i++ {

		patternSlice = append(patternSlice, ogPatternSlice...)

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

func MoreGroupsLeadsToMatch(record []string, pattern []int) bool {

	patternIdentified := GetPattern(record)

	return len(patternIdentified) <= len(pattern)
}

func GenerateArrangements(a *Arrangement, pattern []int) int {

	if !slices.Contains(a.local, "?") {

		if PatternMatch(a.local, pattern) {
			return 1
		} else {
			return 0
		}

	}

	if !MoreGroupsLeadsToMatch(a.local, pattern) {
		return 0
	}

	indexOfChange := slices.Index(a.local, "?")

	a.local[indexOfChange] = "#"
	localChangeDef := []string{}
	localChangeDef = append(localChangeDef, a.local...)

	a.local[indexOfChange] = "."
	localChangeOp := []string{}
	localChangeOp = append(localChangeOp, a.local...)

	return GenerateArrangements(&Arrangement{local: localChangeDef}, pattern) + GenerateArrangements(&Arrangement{local: localChangeOp}, pattern)

}

func arrangementWorker(lines []string, id int, resultChan chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	workerSum := 0

	for _, line := range lines {

		patternUnfolded := ExtractPattern(line, 4)

		record := strings.Split(line, " ")[0]
		recordUnfolded := record

		for i := 0; i < 4; i++ {

			recordUnfolded = recordUnfolded + "?" + record

		}

		lineArrangement := strings.Split(recordUnfolded, "")

		workerSum += GenerateArrangements(&Arrangement{local: lineArrangement}, patternUnfolded)

	}

	fmt.Printf("Woker %d complete\n", id)

	resultChan <- workerSum

}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileLines := strings.Split(fileContent, "\n")

	var wg sync.WaitGroup

	resultChan := make(chan int, 4)

	workerLines := [][]string{}

	workerLines = append(workerLines, fileLines[:len(fileLines)/4])
	workerLines = append(workerLines, fileLines[len(fileLines)/4:len(fileLines)/2])
	workerLines = append(workerLines, fileLines[len(fileLines)/2:len(fileLines)/4*3])
	workerLines = append(workerLines, fileLines[len(fileLines)/4*3:])

	for i := 0; i < 4; i++ {

		wg.Add(1)

		go arrangementWorker(workerLines[i], i, resultChan, &wg)

	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	puzzleAwnser := 0
	for partialSum := range resultChan {
		puzzleAwnser += partialSum
	}

	return puzzleAwnser

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
