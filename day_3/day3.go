package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Serial struct {
	number          int
	didgitPositions []Point
}

type Symbol struct {
	identifier rune
	position   Point
}

type Point struct {
	X int
	Y int
}

func SymbolsOnLine(puzzleLine string, rowNumber int) []Symbol {

	foundSymbolPosisitions := []Symbol{}

	listOfSymbold := []string{"*"}

	// Check if the line has any symbol

	for index, r := range puzzleLine {

		stringConvertedRune := string(r)

		if slices.Contains(listOfSymbold, stringConvertedRune) {

			foundSymbolPosisitions = append(foundSymbolPosisitions, Symbol{identifier: r, position: Point{X: rowNumber, Y: index}})

		}
	}

	return foundSymbolPosisitions
}

func ConvertToNumber(toBeNumber []rune) int {

	exponent := len(toBeNumber) - 1
	returnNumber := 0

	for _, didgit := range toBeNumber {
		returnNumber = returnNumber + ((int(didgit) - 48) * int(math.Pow10(exponent)))
		exponent--
	}

	return returnNumber
}

func GetNumbersAndPositons(puzzleLine string, rowNumber int) []Serial {

	serialsFound := []Serial{}

	// First search for adjecent positions in the first line

	toBeNumber := []rune{}
	didgitPositions := []Point{}

	for position, r := range puzzleLine {

		if 47 < r && r < 58 { // Check if we have a number

			toBeNumber = append(toBeNumber, r)
			didgitPositions = append(didgitPositions, Point{X: rowNumber, Y: position})

		} else {
			if len(toBeNumber) != 0 {

				serialsFound = append(serialsFound, Serial{number: ConvertToNumber(toBeNumber), didgitPositions: didgitPositions})
				toBeNumber = []rune{}
				didgitPositions = []Point{}
			}

		}
	}

	if len(toBeNumber) != 0 {
		serialsFound = append(serialsFound, Serial{number: ConvertToNumber(toBeNumber), didgitPositions: didgitPositions})
	}

	return serialsFound
}

func GetSurroundingPoints(p Point) []Point {

	returnList := []Point{}

	newX := p.X - 1
	newY := p.Y
	if newX > -1 && newY > -1 { // Check if point is valid

		returnList = append(returnList, Point{X: newX, Y: newY})
	}

	newX = p.X - 1
	newY = p.Y + 1
	if newX > -1 && newY > -1 { // Check if point is valid

		returnList = append(returnList, Point{X: newX, Y: newY})
	}

	newX = p.X
	newY = p.Y + 1
	if newX > -1 && newY > -1 { // Check if point is valid

		returnList = append(returnList, Point{X: newX, Y: newY})
	}

	newX = p.X + 1
	newY = p.Y + 1
	if newX > -1 && newY > -1 { // Check if point is valid

		returnList = append(returnList, Point{X: newX, Y: newY})
	}

	newX = p.X + 1
	newY = p.Y
	if newX > -1 && newY > -1 { // Check if point is valid

		returnList = append(returnList, Point{X: newX, Y: newY})
	}

	newX = p.X + 1
	newY = p.Y - 1
	if newX > -1 && newY > -1 { // Check if point is valid

		returnList = append(returnList, Point{X: newX, Y: newY})
	}

	newX = p.X
	newY = p.Y - 1
	if newX > -1 && newY > -1 { // Check if point is valid

		returnList = append(returnList, Point{X: newX, Y: newY})
	}

	newX = p.X - 1
	newY = p.Y - 1
	if newX > -1 && newY > -1 { // Check if point is valid

		returnList = append(returnList, Point{X: newX, Y: newY})
	}

	return returnList
}

func MakeGearRatios(gears []Symbol, serials []Serial) []int {

	gearRatio := 1
	allGearSetsCalulated := []int{}
	gearSet := make(map[int]struct{})

	//symbolLoop:
	for _, symbolLocation := range gears {

		surroundingPoints := GetSurroundingPoints(symbolLocation.position)

		//serialLoop:
		for _, serial := range serials {

			//surroundingPointLoop:
			for _, point := range surroundingPoints {

				//serialDidgitLoop:
				for _, didgitPosition := range serial.didgitPositions {

					if point.X == didgitPosition.X && point.Y == didgitPosition.Y {
						gearSet[serial.number] = struct{}{}
					}

				}

			}

		}

		if len(gearSet) == 2 {
			for key := range gearSet {
				gearRatio = gearRatio * key
			}
			allGearSetsCalulated = append(allGearSetsCalulated, gearRatio)

			gearRatio = 1
		}

		gearSet = make(map[int]struct{})
	}

	return allGearSetsCalulated

}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	puzzleInput := strings.Split(fileContent, "\n")

	gearLocations := []Symbol{}
	serials := []Serial{}

	for rowNumber, row := range puzzleInput {

		gearLocations = append(gearLocations, SymbolsOnLine(row, rowNumber)...)

		serials = append(serials, GetNumbersAndPositons(row, rowNumber)...)
	}

	ratioSum := 0

	gearRatios := MakeGearRatios(gearLocations, serials)

	for _, gearRatio := range gearRatios {

		ratioSum = ratioSum + gearRatio
	}

	return ratioSum
}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
