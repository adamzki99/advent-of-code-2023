package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Card struct {
	cardID          int
	matchingNumbers int
	instances       int
	cardNumbers     []int
	winningNumbers  []int
}

func ConvertSliceToInts(stringSlice []string) []int {

	returnSlice := []int{}

	for _, numberAsString := range stringSlice {

		number, err := strconv.Atoi(numberAsString)

		// The only thing that will trigger the error in this case is a empty string.
		// So there is no need to handle it more than this.

		if err == nil {
			returnSlice = append(returnSlice, number)
		}

	}

	return returnSlice

}

func ExtractCards(puzzleLine string, puzzleRow int) Card {

	extraction := Card{}

	extraction.cardID = puzzleRow + 1

	numbersWithDivider := strings.Split(puzzleLine, ": ")[1]

	numbers := strings.Split(numbersWithDivider, "|")

	cardNumbersAsStrings := strings.Split(numbers[0], " ")
	winningNumbersAsStrings := strings.Split(numbers[1], " ")

	extraction.cardNumbers = ConvertSliceToInts(cardNumbersAsStrings)
	extraction.winningNumbers = ConvertSliceToInts(winningNumbersAsStrings)

	return extraction
}

func GetIntersection(slice1, slice2 []int) []int {

	intersectionSet := make(map[int]bool)

slice1Loop:
	for _, element1 := range slice1 {

		for _, element2 := range slice2 {

			if element1 == element2 {
				intersectionSet[element1] = true
				continue slice1Loop
			}
		}
	}

	// Now turn the set into a sorted slice

	returnSlice := []int{}

	for key := range intersectionSet {

		returnSlice = append(returnSlice, key)

	}

	sort.Ints(returnSlice)

	return returnSlice

}

func GenerateCardCopies(cards *[]Card, currentCard Card) {

	for m := 0; m < currentCard.instances; m++ {
		for i := currentCard.cardID; i < currentCard.cardID+currentCard.matchingNumbers; i++ {

			(*cards)[i].instances = 1 + (*cards)[i].instances

		}
	}

}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	puzzleLines := strings.Split(fileContent, "\n")

	cards := []Card{}

	for index, line := range puzzleLines {

		currentCard := ExtractCards(line, index)
		currentCard.instances = 1
		winningMatches := GetIntersection(currentCard.cardNumbers, currentCard.winningNumbers)
		currentCard.matchingNumbers = len(winningMatches)

		cards = append(cards, currentCard)
	}

	for _, card := range cards {

		GenerateCardCopies(&cards, card)
	}

	// Sum together all instances

	instanceSum := 0

	for _, card := range cards {

		instanceSum = instanceSum + card.instances

	}

	return instanceSum
}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
