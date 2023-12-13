package main

import (
	"os"
	"testing"

	"github.com/adamzki99/advent-of-code-2023/packages/testhelp"
)

func TestGetPattern(t *testing.T) {

	result := GetPattern([]string{"#", ".", "#", ".", "#", "#", "#"})
	expected := []int{1, 1, 3}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("GetPattern function test failed.")
	}

}

func TestExtractPattern(t *testing.T) {

	result := ExtractPattern("#.#.### 1,1,3", 4)
	expected := []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("ExtractPattern function test failed.")
	}

	result = ExtractPattern("?#?#?#?#?#?#?#? 1,3,1,6", 4)
	expected = []int{1, 3, 1, 6, 1, 3, 1, 6, 1, 3, 1, 6, 1, 3, 1, 6, 1, 3, 1, 6}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("ExtractPattern function test failed.")
	}

	result = ExtractPattern("??#??#???#????????? 10,2,1,1", 4)
	expected = []int{10, 2, 1, 1, 10, 2, 1, 1, 10, 2, 1, 1, 10, 2, 1, 1, 10, 2, 1, 1}

	if !testhelp.AreIntSlicesEqual(result, expected) {
		t.Errorf("ExtractPattern function test failed.")
	}

}

func TestPatternMatch(t *testing.T) {

	result := PatternMatch([]string{"#", ".", "#", ".", "#", "#", "#"}, []int{1, 1, 3})
	expected := true
	if result != expected {
		t.Errorf("PatternMatch function test failed. Expected: %t, Got: %t", expected, result)
	}

}

func TestGenerateArrangements(t *testing.T) {

	input := &Arrangement{local: []string{"?", ".", "#"}}

	result := GenerateArrangements(input, []int{1, 1})
	expected := 1
	if result != expected {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expected, result)
	}

	//

	input = &Arrangement{local: []string{"?", "?", "?", ".", "#", "#", "#"}}
	result = GenerateArrangements(input, []int{1, 1, 3})
	expected = 1
	if result != expected {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expected, result)
	}

	input = &Arrangement{local: []string{".", "?", "?", ".", ".", "?", "?", ".", ".", ".", "?", "#", "#", "."}}
	result = GenerateArrangements(input, []int{1, 1, 3})
	expected = 4
	if result != expected {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expected, result)
	}

	input = &Arrangement{local: []string{"?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?"}}
	result = GenerateArrangements(input, []int{1, 3, 1, 6})
	expected = 1
	if result != expected {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expected, result)
	}

	input = &Arrangement{local: []string{"?", "?", "?", "?", ".", "#", ".", ".", ".", "#", ".", ".", "."}}
	result = GenerateArrangements(input, []int{4, 1, 1})
	expected = 1
	if result != expected {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expected, result)
	}

	input = &Arrangement{local: []string{"?", "?", "?", "?", ".", "#", "#", "#", "#", "#", "#", ".", ".", "#", "#", "#", "#", "#", "."}}
	result = GenerateArrangements(input, []int{1, 6, 5})
	expected = 4
	if result != expected {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expected, result)
	}

	input = &Arrangement{local: []string{"?", "#", "#", "#", "?", "?", "?", "?", "?", "?", "?", "?"}}
	result = GenerateArrangements(input, []int{3, 2, 1})
	expected = 10
	if result != expected {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expected, result)
	}

	input = &Arrangement{local: []string{"?", "?", "#", "?", "?", "#", "?", "?", "?", "#", "?", "?", "?", "?", "?", "?", "?", "?", "?"}}
	result = GenerateArrangements(input, []int{10, 2, 1, 1})
	expected = 15
	if result != expected {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`)

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
	expected := 525152

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
