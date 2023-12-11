package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Position struct {
	col int
	row int
}

// Returns true if there is no galaxy in the column
func InspectColumn(image *[]string, columnIndex, lineIndex int) bool {

	if lineIndex == len(*image) {
		return true
	}

	if (*image)[lineIndex][columnIndex] == '#' {
		return false
	}

	return InspectColumn(image, columnIndex, lineIndex+1)

}

// Here columnIndex is given as where the search starts
func InspectRow(imageLine *string, columnIndex int) bool {

	if columnIndex == len(*imageLine) {
		return true
	}

	if (*imageLine)[columnIndex] == '#' {
		return false
	}

	return InspectRow(imageLine, columnIndex+1)

}

// image: A slice of lines representing the image
func ExpandImage(image *[]string) {

	// Lets start by finding empty columns

	emptyColumIndexes := []int{}

	for i := 0; i < len((*image)[0]); i++ {

		if InspectColumn(image, i, 0) {
			emptyColumIndexes = append(emptyColumIndexes, i)
		}

	}

	emptyRowIndexes := []int{}

	for i, row := range *image {

		if InspectRow(&row, 0) {
			emptyRowIndexes = append(emptyRowIndexes, i)
		}
	}

	// Now we can add the space

	for c, eColID := range emptyColumIndexes {

		for i, row := range *image {

			(*image)[i] = row[:eColID+c] + "." + row[eColID+c:]

		}

	}

	// lets create a "empty space" row string

	emptySpaceRow := ""
	for i := 0; i < len((*image)[0]); i++ {

		emptySpaceRow = emptySpaceRow + "."

	}

	for i, eRowID := range emptyRowIndexes {

		*image = slices.Replace(*image, eRowID+i, eRowID+i, emptySpaceRow)

	}

}

// Returns highest galaxy index
func AssignNumbersToGalaxies(pixels *[][]string) []Position {

	galaxyIndex := 0

	galaxyLocations := []Position{}

	for i := range *pixels {

		for j := range (*pixels)[i] {

			if (*pixels)[i][j] == "#" {

				(*pixels)[i][j] = strconv.Itoa(galaxyIndex)
				galaxyIndex++

				galaxyLocations = append(galaxyLocations, Position{col: i, row: j})
			}
		}

	}

	return galaxyLocations
}

func Distance(start, goal Position) int {

	manhattanDistance := math.Abs(float64(start.row)-float64(goal.row)) + math.Abs(float64(start.col)-float64(goal.col))

	return int(manhattanDistance)

}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	image := strings.Split(fileContent, "\n")

	ExpandImage(&image)

	pixels := [][]string{}

	for _, row := range image {
		pixels = append(pixels, strings.Split(row, ""))
	}

	galaxyLocations := AssignNumbersToGalaxies(&pixels)

	distanceSum := 0

	for _, g1 := range galaxyLocations {

		for _, g2 := range galaxyLocations {

			distanceSum = distanceSum + Distance(g1, g2)

		}

	}

	return distanceSum / 2

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
