package main

import (
	"os"
	"testing"
)

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

func TestStringOfNumbersToSingle(t *testing.T) {

	result := StringOfNumbersToSingle("Test:   88 18 7", " ")
	expected := 88187

	if result != expected {
		t.Errorf("StringOfNumbersToSingle function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = StringOfNumbersToSingle("Test:   0 69 1", " ")
	expected = 691

	if result != expected {
		t.Errorf("StringOfNumbersToSingle function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = StringOfNumbersToSingle("Test:   52 50 48", " ")
	expected = 525048

	if result != expected {
		t.Errorf("StringOfNumbersToSingle function test failed. Expected: %d, Got: %d", expected, result)
	}

}

func TestExtractRace(t *testing.T) {

	fileLines := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	result := ExtractRace(&fileLines)
	expected := Race{
		time:     71530,
		distance: 940200,
	}

	if result != expected {
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
	expected := 71503

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
