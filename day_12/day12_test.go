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

	GenerateArrangements(input, []int{1, 1})
	expected := &Arrangement{
		local:        []string{"?", ".", "#"},
		damaged:      &Arrangement{local: []string{"#", ".", "#"}},
		operational:  &Arrangement{local: []string{".", ".", "#"}},
		matchesUnder: 1,
	}

	if !testhelp.AreStringSlicesEqual(input.local, expected.local) ||
		!testhelp.AreStringSlicesEqual(input.damaged.local, expected.damaged.local) ||
		!testhelp.AreStringSlicesEqual(input.operational.local, expected.operational.local) ||
		input.matchesUnder != expected.matchesUnder {
		t.Errorf("GenerateArrangements function test failed.")
	}

	//

	input = &Arrangement{local: []string{"?", "?", "?", ".", "#", "#", "#"}}
	GenerateArrangements(input, []int{1, 1, 3})
	expectedMatches := 1
	if input.matchesUnder != expectedMatches {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expectedMatches, input.matchesUnder)
	}

	input = &Arrangement{local: []string{".", "?", "?", ".", ".", "?", "?", ".", ".", ".", "?", "#", "#", "."}}
	GenerateArrangements(input, []int{1, 1, 3})
	expectedMatches = 4
	if input.matchesUnder != expectedMatches {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expectedMatches, input.matchesUnder)
	}

	input = &Arrangement{local: []string{"?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?"}}
	GenerateArrangements(input, []int{1, 3, 1, 6})
	expectedMatches = 1
	if input.matchesUnder != expectedMatches {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expectedMatches, input.matchesUnder)
	}

	input = &Arrangement{local: []string{"?", "?", "?", "?", ".", "#", ".", ".", ".", "#", ".", ".", "."}}
	GenerateArrangements(input, []int{4, 1, 1})
	expectedMatches = 1
	if input.matchesUnder != expectedMatches {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expectedMatches, input.matchesUnder)
	}

	input = &Arrangement{local: []string{"?", "?", "?", "?", ".", "#", "#", "#", "#", "#", "#", ".", ".", "#", "#", "#", "#", "#", "."}}
	GenerateArrangements(input, []int{1, 6, 5})
	expectedMatches = 4
	if input.matchesUnder != expectedMatches {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expectedMatches, input.matchesUnder)
	}

	input = &Arrangement{local: []string{"?", "#", "#", "#", "?", "?", "?", "?", "?", "?", "?", "?"}}
	GenerateArrangements(input, []int{3, 2, 1})
	expectedMatches = 10
	if input.matchesUnder != expectedMatches {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expectedMatches, input.matchesUnder)
	}

	input = &Arrangement{local: []string{"?", "?", "#", "?", "?", "#", "?", "?", "?", "#", "?", "?", "?", "?", "?", "?", "?", "?", "?"}}
	GenerateArrangements(input, []int{10, 2, 1, 1})
	expectedMatches = 15
	if input.matchesUnder != expectedMatches {
		t.Errorf("GenerateArrangements function test failed. Expected: %d, Got: %d", expectedMatches, input.matchesUnder)
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
