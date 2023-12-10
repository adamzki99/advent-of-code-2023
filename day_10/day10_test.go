package main

import (
	"os"
	"strings"
	"testing"
)

func TestNextPipeIndex(t *testing.T) {

	result := NextPipeIndex("-", 6, 5, &Momentum{horisontalDir: -1, verticalDir: 0})
	expected := 5
	if result != expected {
		t.Errorf("NextPipeIndex function test failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestStepFactory(t *testing.T) {

	m := Momentum{horisontalDir: 0, verticalDir: -1}
	input := strings.Split(".F-7.", "")
	stepFactory := StepFactory(input, 1, 5, &m)
	result := stepFactory()
	expected := "-"
	if result != expected {
		t.Errorf("StepFactory function test failed. Expected: %s, Got: %s", expected, result)
	}
	result = stepFactory()
	expected = "7"
	if result != expected {
		t.Errorf("StepFactory function test failed. Expected: %s, Got: %s", expected, result)
	}

	m = Momentum{horisontalDir: 1, verticalDir: 0}
	input = strings.Split(".F-7..|.|.", "")
	stepFactory = StepFactory(input, 2, 5, &m)
	result = stepFactory()
	expected = "7"
	if result != expected {
		t.Errorf("StepFactory function test failed. Expected: %s, Got: %s", expected, result)
	}
	result = stepFactory()
	expected = "|"
	if result != expected {
		t.Errorf("StepFactory function test failed. Expected: %s, Got: %s", expected, result)
	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`)

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
	expected := 8

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
