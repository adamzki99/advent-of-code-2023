package main

import (
	"os"
	"testing"
)

func areRaceSlicesEqual(rs1, rs2 []Race) bool {

	if len(rs1) != len(rs2) {
		return false
	}

	for i := range rs1 {
		if rs1[i] != rs2[i] {
			return false
		}
	}

	return true
}

func areIntSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {

		if slice1[i] != slice2[i] {
			return false
		}

	}
	return true
}

func TestGetDistanceTraveled(t *testing.T) {

	result := GetDistanceTraveled(0, 7)
	expected := 0
	if result != expected {
		t.Errorf("GetDistanceTraveled function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = GetDistanceTraveled(1, 7)
	expected = 6
	if result != expected {
		t.Errorf("GetDistanceTraveled function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = GetDistanceTraveled(2, 7)
	expected = 10
	if result != expected {
		t.Errorf("GetDistanceTraveled function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = GetDistanceTraveled(3, 7)
	expected = 12
	if result != expected {
		t.Errorf("GetDistanceTraveled function test failed. Expected: %d, Got: %d", expected, result)
	}

}

func TestStringOfNumbersToSliceOfNumbers(t *testing.T) {

	result := StringOfNumbersToSliceOfNumbers("88 18 7", " ")
	expected := []int{88, 18, 7}

	if !areIntSlicesEqual(result, expected) {
		t.Error("StringOfNumbersToSliceOfNumbers function test 1 failed.")
	}

	result = StringOfNumbersToSliceOfNumbers("0 69 1", " ")
	expected = []int{0, 69, 1}

	if !areIntSlicesEqual(result, expected) {
		t.Error("StringOfNumbersToSliceOfNumbers function test 2 failed.")
	}

	result = StringOfNumbersToSliceOfNumbers("52 50 48", " ")
	expected = []int{52, 50, 48}

	if !areIntSlicesEqual(result, expected) {
		t.Error("StringOfNumbersToSliceOfNumbers function test 3 failed.")
	}

}

func TestExtractRaces(t *testing.T) {

	fileLines := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	result := ExtractRaces(&fileLines)
	expected := []Race{
		{
			time:     7,
			distance: 9,
		},
		{
			time:     15,
			distance: 40,
		},
		{
			time:     30,
			distance: 200,
		},
	}

	if !areRaceSlicesEqual(result, expected) {
		t.Errorf("ExtractRaces function test failed.")

	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`Time:      7  15   30
Distance:  9  40  200`,
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
	expected := 288

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
