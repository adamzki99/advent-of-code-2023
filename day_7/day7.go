package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Hand struct {
	bid     int
	hType   int // Five of a kind = 7, four of a kind = 6, ..., high card = 1
	hString string
	//label    int // A = 13, K = 12, Q = 11, ..., 2 = 1
}

func LabelConversion(label string) int {

	switch label {
	case "A":
		return 41
	case "K":
		return 37
	case "Q":
		return 31
	case "T":
		return 23
	case "9":
		return 19
	case "8":
		return 17
	case "7":
		return 13
	case "6":
		return 11
	case "5":
		return 7
	case "4":
		return 5
	case "3":
		return 3
	case "2":
		return 2
	case "J":
		return 0
	default:
		return -1
	}
}

func ExtractOccurrences(hand []string) []int {

	occurrencesByLabel := make(map[int]int)

	for _, element := range hand {
		occurrencesByLabel[LabelConversion(element)]++
	}

	occurrences := []int{}

	for _, v := range occurrencesByLabel {

		occurrences = append(occurrences, v)
	}

	sort.Ints(occurrences)
	slices.Reverse(occurrences)

	return occurrences
}


func HandTypeExtraction(hand string) int {

	occurrences := ExtractOccurrences(strings.Split(hand, ""))

	if len(occurrences) > 1 && occurrences[0] == 3 && occurrences[1] == 2 {
		return 61
	} else if len(occurrences) > 1 && occurrences[0] == 2 && occurrences[1] == 2 {
		return 53
	} else if occurrences[0] == 5 {
		return 71
	} else if occurrences[0] == 4 {
		return 67
	} else if occurrences[0] == 3 {
		return 59
	} else if occurrences[0] == 2 {
		return 47
	}

	return 43
}

func ConstructString(s []string) string {

	returnString := ""

	for _, v := range s {
		returnString = returnString + v
	}

	return returnString
}

// Find all possible cobinations of placing a joker, and return the best type
func HandTypeExtractionJoker(hand string) int {

	if strings.Count(hand, "J") == 0 {
		return HandTypeExtraction(hand)
	}

	bestHand := 0

	symbols := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

	handSymbols := strings.Split(hand, "")
	jokerIndexes := []int{}

	for i, hS := range handSymbols {
		if hS == "J" {
			jokerIndexes = append(jokerIndexes, i)
		}
	}

	for _, jokerIndex := range jokerIndexes {

		for _, symbol := range symbols {

			handSymbols[jokerIndex] = symbol

			stringToSend := ConstructString(handSymbols)
			currentHand := HandTypeExtractionJoker(stringToSend)

			if bestHand < currentHand {
				bestHand = currentHand
			}

		}

		handSymbols[jokerIndex] = "J"
	}

	return bestHand
}

// Usage is similar to "Is rank of h1 higher than rank of h2?"
func CompareHands(h1, h2 Hand) bool {

	if h1.hType < h2.hType {
		return false
	}

	if h1.hType > h2.hType {
		return true
	}

	h1Chars := strings.Split(h1.hString, "")
	h2Chars := strings.Split(h2.hString, "")

	for i := 0; i < 5; i++ { //Hands are only 5 characters long

		if LabelConversion(h1Chars[i]) == LabelConversion(h2Chars[i]) {
			continue
		}

		if LabelConversion(h1Chars[i]) > LabelConversion(h2Chars[i]) {
			return true
		} else {
			return false
		}

	}

	return false
}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileLines := strings.Split(fileContent, "\n")

	fmt.Printf("Lines read: %d\n", len(fileLines))

	hands := []Hand{}

	for _, line := range fileLines {

		puzzleInput := strings.Split(line, " ")

		currentHand := Hand{}

		currentHand.bid, err = strconv.Atoi(puzzleInput[1])

		if err != nil {
			fmt.Println(err)
			return -1
		}

		if strings.Count(puzzleInput[0], "J") == 0 {
			currentHand.hType = HandTypeExtraction(puzzleInput[0])
		} else {
			currentHand.hType = HandTypeExtractionJoker(puzzleInput[0])
		}
		currentHand.hString = puzzleInput[0]
		hands = append(hands, currentHand)
	}

	// Now ranking the hands

	for iPass := 1; iPass < len(hands); iPass++ {
		for i := 0; i < len(hands)-iPass; i++ {

			if CompareHands(hands[i], hands[i+1]) {
				temp := hands[i]
				hands[i] = hands[i+1]
				hands[i+1] = temp
			}
		}
	}

	puzzleAwnser := 0

	for i, hand := range hands {
		puzzleAwnser = puzzleAwnser + (hand.bid * (i + 1))
	}

	return puzzleAwnser

}

func main() {
	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
