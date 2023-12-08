package main

import (
	"fmt"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Place struct {
	current  string
	left     *Place
	right    *Place
	treeRoot *Place
}

func (p *Place) AppendLeft(place string, tr *Place) {

	if p.current == place {
		return
	}

	existingPlace := p.treeRoot.SearchPlaces(place, &[]string{})

	if existingPlace != nil {
		p.left = existingPlace
		return
	}

	p.left = &Place{current: place, treeRoot: p.treeRoot}

}

func (p *Place) AppendRight(place string, tr *Place) {

	if p.current == place {
		return
	}

	existingPlace := p.treeRoot.SearchPlaces(place, &[]string{})

	if existingPlace != nil {
		p.right = existingPlace
		return
	}

	p.right = &Place{current: place, treeRoot: p.treeRoot}

}

func sliceContains(s *[]string, e string) bool {
	for _, v := range *s {
		if v == e {
			return true
		}
	}

	return false
}

func (p *Place) SearchPlaces(placeToFind string, visitedPlaces *[]string) *Place {

	var returnPlace *Place

	if p == nil {
		return returnPlace
	}

	*visitedPlaces = append(*visitedPlaces, p.current)

	if p.current == placeToFind {
		return p
	}

	if p.left != nil && !sliceContains(visitedPlaces, p.left.current) {
		returnPlace = p.left.SearchPlaces(placeToFind, visitedPlaces)
	}

	if returnPlace == nil && p.right != nil && !sliceContains(visitedPlaces, p.right.current) {
		returnPlace = p.right.SearchPlaces(placeToFind, visitedPlaces)
	}

	return returnPlace

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

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	puzzleLines := strings.Split(fileContent, "\n")

	nodeLines := puzzleLines[2:]
	puzzleLineFactory := LineFactory(nodeLines)

	currentLine, _ := puzzleLineFactory()

	p1, p2, p3 := SplitNodeLine(currentLine)

	start := Place{current: p1}
	start.treeRoot = &start
	start.AppendLeft(p2, &start)
	start.AppendRight(p3, &start)

	nodeLines = removeElement(nodeLines, 0)
	puzzleLineFactory = LineFactory(nodeLines)

	for {

		currentLine, i := puzzleLineFactory()

		p1, p2, p3 = SplitNodeLine(currentLine)

		currentNode := start.SearchPlaces(p1, &[]string{})

		if currentNode == nil {
			continue
		}

		currentNode.AppendLeft(p2, &start)
		currentNode.AppendRight(p3, &start)

		nodeLines = removeElement(nodeLines, i)
		puzzleLineFactory = LineFactory(nodeLines)

		if len(nodeLines) == 0 {
			break
		}

	}

	stepFactory := CharFactory(puzzleLines[0])
	steps := 0

	currentPlace := &start

	for {

		step := stepFactory()
		steps++

		if step == "L" {
			currentPlace = currentPlace.left
		} else {
			currentPlace = currentPlace.right
		}

		if currentPlace.current == "ZZZ" {
			break
		}

	}

	return steps

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
