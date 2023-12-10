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
	r1, r2 := stepFactory()
	e1 := "-"
	e2 := 2
	if r1 != e1 || r2 != e2 {
		t.Errorf("StepFactory function test failed.")
	}
	r1, r2 = stepFactory()
	e1 = "7"
	e2 = 3
	if r1 != e1 || r2 != e2 {
		t.Errorf("StepFactory function test failed.")
	}

	m = Momentum{horisontalDir: 1, verticalDir: 0}
	input = strings.Split(".F-7..|.|.", "")
	stepFactory = StepFactory(input, 2, 5, &m)
	r1, r2 = stepFactory()
	e1 = "7"
	e2 = 3
	if r1 != e1 || r2 != e2 {
		t.Errorf("StepFactory function test failed.")
	}
	r1, r2 = stepFactory()
	e1 = "|"
	e2 = 8
	if r1 != e1 || r2 != e2 {
		t.Errorf("StepFactory function test failed.")
	}

}

func TestDoesPointTouchAnyPointInSet(t *testing.T) {

	result := DoesPointTouchAnyPointInSet(1, 5, []int{2, 3, 4})
	expected := true
	if result != expected {
		t.Errorf("DoesPointTouchAnyPointInSet function test failed. Expected: %t, Got: %t", expected, result)
	}

	result = DoesPointTouchAnyPointInSet(1, 5, []int{2, 6, 4})
	expected = true
	if result != expected {
		t.Errorf("DoesPointTouchAnyPointInSet function test failed. Expected: %t, Got: %t", expected, result)
	}

	result = DoesPointTouchAnyPointInSet(2, 5, []int{2, 6, 4})
	expected = false
	if result != expected {
		t.Errorf("DoesPointTouchAnyPointInSet function test failed. Expected: %t, Got: %t", expected, result)
	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`)

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
	expected := 4

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
