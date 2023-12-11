package main

import (
	"os"
	"testing"

	"github.com/adamzki99/advent-of-code-2023/packages/testhelp"
)

func TestInspectColumn(t *testing.T) {

	input := []string{"...#......"}
	result := InspectColumn(&input, 2, 0)
	expected := true
	if result != expected {
		t.Errorf("InspectColumn function test failed. Expected: %t, Got: %t", expected, result)
	}

	input = []string{
		"...#......",
		"..........",
	}
	result = InspectColumn(&input, 2, 0)
	expected = true
	if result != expected {
		t.Errorf("InspectColumn function test failed. Expected: %t, Got: %t", expected, result)
	}

	input = []string{
		"...#......",
		"..........",
		".......#..",
	}
	result = InspectColumn(&input, 7, 0)
	expected = false
	if result != expected {
		t.Errorf("InspectColumn function test failed. Expected: %t, Got: %t", expected, result)
	}

}

func TestInspectRow(t *testing.T) {

	input := "...#......"
	result := InspectRow(&input, 0)
	expected := false
	if result != expected {
		t.Errorf("InspectRow function test failed. Expected: %t, Got: %t", expected, result)
	}

	input = ".........."
	result = InspectRow(&input, 0)
	expected = true
	if result != expected {
		t.Errorf("InspectRow function test failed. Expected: %t, Got: %t", expected, result)
	}

}

func TestExpandImage(t *testing.T) {

	input := []string{
		"...#.",
		".....",
		"#....",
		".....",
		".....",
	}
	ExpandImage(&input)
	expected := []string{
		".....#..",
		"........",
		"........",
		"#.......",
		"........",
		"........",
		"........",
		"........",
	}

	if !testhelp.AreStringSlicesEqual(input, expected) {
		t.Errorf("ExpandImage function test 1 failed.")
	}

	input = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	ExpandImage(&input)
	expected = []string{
		"....#........",
		".........#...",
		"#............",
		".............",
		".............",
		"........#....",
		".#...........",
		"............#",
		".............",
		".............",
		".........#...",
		"#....#.......",
	}

	if !testhelp.AreStringSlicesEqual(input, expected) {
		t.Errorf("ExpandImage function test 2 failed.")
	}
}

func TestAssignNumbersToGalaxies(t *testing.T) {

	input := [][]string{
		{".", "#"},
		{"#", "."},
	}

	AssignNumbersToGalaxies(&input)
	expected := [][]string{
		{".", "0"},
		{"1", "."},
	}

	for i := range input {
		if !testhelp.AreStringSlicesEqual(input[i], expected[i]) {
			t.Errorf("AssignNumbersToGalaxies function test failed.")
		}
	}

}

func TestDistance(t *testing.T) {

	result := Distance(Position{6, 1}, Position{11, 5})
	expected := 9
	if result != expected {
		t.Errorf("AStarDistance function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = Distance(Position{4, 0}, Position{10, 9})
	expected = 15
	if result != expected {
		t.Errorf("AStarDistance function test failed. Expected: %d, Got: %d", expected, result)
	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`)

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
	expected := 374

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
