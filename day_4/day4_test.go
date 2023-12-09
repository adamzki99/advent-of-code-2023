package main

import (
	"os"
	"testing"

	"github.com/adamzki99/advent-of-code-2023/packages/testhelp"
)

func areCardsEqual(card1, card2 Card) bool {

	if card1.cardID != card2.cardID {
		return false
	}

	if card1.matchingNumbers != card2.matchingNumbers {
		return false
	}

	if card1.instances != card2.instances {
		return false
	}

	if !testhelp.AreIntSlicesEqual(card1.cardNumbers, card2.cardNumbers) {
		return false
	}

	if !testhelp.AreIntSlicesEqual(card1.winningNumbers, card2.winningNumbers) {
		return false
	}

	return true

}

func areCardSlicesEqual(slice1, slice2 []Card) bool {

	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {

		if !areCardsEqual(slice1[i], slice2[i]) {
			return false
		}
	}

	return true
}

func TestConvertSliceToInts(t *testing.T) {

	result := ConvertSliceToInts([]string{"41", "48", "83", "86", "17"})
	expected := []int{41, 48, 83, 86, 17}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("ConvertSliceToInts function test failed.")
	}

}

func TestExtractCards(t *testing.T) {

	result := ExtractCards("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 0)
	expected := Card{
		cardID:         1,
		cardNumbers:    []int{41, 48, 83, 86, 17},
		winningNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	if !areCardsEqual(result, expected) {
		t.Errorf("ExtractCards function test 1 failed.")
	}

	result = ExtractCards("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2)
	expected = Card{
		cardID:         3,
		cardNumbers:    []int{1, 21, 53, 59, 44},
		winningNumbers: []int{69, 82, 63, 72, 16, 21, 14, 1},
	}

	if !areCardsEqual(result, expected) {
		t.Errorf("ExtractCards function test 2 failed.")
	}
}

func TestGetIntersection(t *testing.T) {

	result := GetIntersection([]int{}, []int{})
	expected := []int{}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("GetIntersection function test 1 failed.")
	}

	result = GetIntersection([]int{1}, []int{})
	expected = []int{}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("GetIntersection function test 1 failed.")
	}

	result = GetIntersection([]int{1, 2, 9, 10}, []int{9, 2, 1, 9})
	expected = []int{1, 2, 9}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("GetIntersection function test 1 failed.")
	}

}

func TestGenerateCardCopies(t *testing.T) {

	cards := []Card{
		{
			cardID:         1,
			instances:      1,
			cardNumbers:    []int{41, 48, 83, 86, 17},
			winningNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			cardID:         2,
			instances:      1,
			cardNumbers:    []int{13, 32, 20, 16, 61},
			winningNumbers: []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			cardID:         3,
			instances:      1,
			cardNumbers:    []int{1, 21, 53, 59, 44},
			winningNumbers: []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			cardID:         4,
			instances:      1,
			cardNumbers:    []int{41, 92, 73, 84, 69},
			winningNumbers: []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			cardID:         5,
			instances:      1,
			cardNumbers:    []int{87, 83, 26, 28, 32},
			winningNumbers: []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			cardID:         6,
			instances:      1,
			cardNumbers:    []int{31, 18, 13, 56, 72},
			winningNumbers: []int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}

	GenerateCardCopies(&cards, Card{cardID: 1, matchingNumbers: 4, instances: 1, cardNumbers: []int{41, 48, 83, 86, 17}, winningNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}})
	expected := []Card{
		{
			cardID:         1,
			instances:      1,
			cardNumbers:    []int{41, 48, 83, 86, 17},
			winningNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			cardID:         2,
			instances:      2,
			cardNumbers:    []int{13, 32, 20, 16, 61},
			winningNumbers: []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			cardID:         3,
			instances:      2,
			cardNumbers:    []int{1, 21, 53, 59, 44},
			winningNumbers: []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			cardID:         4,
			instances:      2,
			cardNumbers:    []int{41, 92, 73, 84, 69},
			winningNumbers: []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			cardID:         5,
			instances:      2,
			cardNumbers:    []int{87, 83, 26, 28, 32},
			winningNumbers: []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			cardID:         6,
			instances:      1,
			cardNumbers:    []int{31, 18, 13, 56, 72},
			winningNumbers: []int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}

	if !areCardSlicesEqual(cards, expected) {
		t.Errorf("GenerateCardCopies function test failed.")
	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
	)

	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}

	// Remove temporary test file with the end of the test
	defer os.Remove(tmpfile.Name())

	// Check write contect to file
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	result := SolvePuzzle(tmpfile.Name())
	expected := 30

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
