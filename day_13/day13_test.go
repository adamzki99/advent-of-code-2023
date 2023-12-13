package main

import (
	"os"
	"testing"

	"github.com/adamzki99/advent-of-code-2023/packages/testhelp"
)

func TestFindColumnWithReflection(t *testing.T) {

	input := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}

	result := FindColumnWithReflection(&input)
	expected := []int{5}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("FindColumnWithReflection function test failed. Expected %v, Got %v", expected, result)
	}

}

func TestFindRowWithReflection(t *testing.T) {

	input := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}

	result := FindRowWithReflection(&input)
	expected := []int{3}
	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("FindRowWithReflection function test failed. Expected %v, Got %v", expected, result)
	}

	input = []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}

	result = FindRowWithReflection(&input)
	expected = []int{4}
	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("FindRowWithReflection function test failed. Expected %v, Got %v", expected, result)
	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(``)

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
	expected := -1

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
