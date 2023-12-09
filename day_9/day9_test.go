package main

import (
	"os"
	"testing"
)

func areIntSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {

		if v != slice2[i] {
			return false
		}

	}
	return true
}

func TestProduceDifferences(t *testing.T) {

	input := []int{21, 15, 10, 6, 3, 1}
	result := ProduceDifferences(&input)
	expected := []int{6, 5, 4, 3, 2}
	if !areIntSlicesEqual(result, expected) {
		t.Error("ProduceDifferences function test failed.")
	}

	input = []int{1, 1, 1, 1, 1}
	result = ProduceDifferences(&input)
	expected = []int{0, 0, 0, 0}
	if !areIntSlicesEqual(result, expected) {
		t.Error("ProduceDifferences function test failed.")
	}

}

func TestStringLineToIntSlice(t *testing.T) {

	input := "1   3   6  10  15  21"
	result := StringLineToIntSlice(&input)
	expected := []int{1, 3, 6, 10, 15, 21}
	if !areIntSlicesEqual(result, expected) {
		t.Error("StringLineToIntSlice function test failed.")
	}

}

func TestCalculateNextValue(t *testing.T) {

	l1 := []int{0, 3, 6, 9, 12, 15}
	l2 := []int{3, 3, 3, 3, 3, 3}
	CalculateNextValue(&l1, &l2)
	expected := []int{0, 3, 6, 9, 12, 15, 12}
	if !areIntSlicesEqual(l1, expected) {
		t.Error("CalculateNextValue function test failed.")
	}

	l1 = []int{3, 3, 3, 3, 3}
	l2 = []int{0, 0, 0, 0}
	CalculateNextValue(&l1, &l2)
	expected = []int{3, 3, 3, 3, 3, 3}
	if !areIntSlicesEqual(l1, expected) {
		t.Error("CalculateNextValue function test failed.")
	}

}

func TestAllElementsAreZero(t *testing.T) {

	input := []int{3, 3, 3, 3, 3}
	result := AllElementsAreZero(&input)
	expected := false
	if result != expected {
		t.Errorf("AllElementsAreZero function test failed. Expected: %t, Got: %t", expected, result)
	}

	input = []int{0, 0, 0, 0}
	result = AllElementsAreZero(&input)
	expected = true
	if result != expected {
		t.Errorf("AllElementsAreZero function test failed. Expected: %t, Got: %t", expected, result)
	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`)

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
	expected := 2

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
