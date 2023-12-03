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

type Point struct {
	X int
	Y int
}

func SymbolsOnLine(puzzleLine string, rowNumber int) []Point {

	foundSymbolPosisitions := []Point{}

	listOfSymbold := []string{".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// Check if the line has any symbol

	for index, r := range puzzleLine {

		stringConvertedRune := string(r)

		if !slices.Contains(listOfSymbold, stringConvertedRune) {

			foundSymbolPosisitions = append(foundSymbolPosisitions, Point{X: rowNumber, Y: index})

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

func IsSerial(serialToInspect Serial, symbolLocations []Point) bool {

	for _, symbolLocation := range symbolLocations {

		surroundingPoints := GetSurroundingPoints(symbolLocation)

		for _, point := range surroundingPoints {

			for _, didgitPosition := range serialToInspect.didgitPositions {

				if point.X == didgitPosition.X && point.Y == didgitPosition.Y {
					return true
				}

			}

		}

	}

	return false
}

func RunProgram(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	puzzleInput := strings.Split(fileContent, "\n")

	symbolLocation := []Point{}
	serials := []Serial{}

	for rowNumber, row := range puzzleInput {

		symbolLocation = append(symbolLocation, SymbolsOnLine(row, rowNumber)...)

		serials = append(serials, GetNumbersAndPositons(row, rowNumber)...)
	}

	serialSum := 0

	for _, serial := range serials {

		if IsSerial(serial, symbolLocation) {
			serialSum = serialSum + serial.number

		} 
	}

	return serialSum
}

func main() {
	fmt.Println(RunProgram("puzzle_input.txt"))
}
