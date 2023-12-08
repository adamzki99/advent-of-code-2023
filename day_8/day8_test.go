package main

import (
	"os"
	"testing"
)

func TestCharFactory(t *testing.T) {

	input := "AB12"
	factory := CharFactory(input)

	result := factory()
	expected := "A"
	if result != expected {
		t.Errorf("CharFactory function test failed. Expected: %s, Got: %s", expected, result)
	}

	result = factory()
	expected = "B"
	if result != expected {
		t.Errorf("CharFactory function test failed. Expected: %s, Got: %s", expected, result)
	}

	result = factory()
	expected = "1"
	if result != expected {
		t.Errorf("CharFactory function test failed. Expected: %s, Got: %s", expected, result)
	}

	result = factory()
	expected = "2"
	if result != expected {
		t.Errorf("CharFactory function test failed. Expected: %s, Got: %s", expected, result)
	}

}

func areStringSlicesEqual(slice1, slice2 []string) bool {
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

func TestRebuildSliceWithIndexes(t *testing.T) {

	input := []string{"a", "b", "c", "d"}
	result := RebuildSliceWithIndexes(input, []int{1})
	expected := []string{"b"}

	if !areStringSlicesEqual(result, expected) {
		t.Error("RebuildSliceWithIndexes function test failed.")
	}

	input = []string{"a", "b", "c", "d"}
	result = RebuildSliceWithIndexes(input, []int{1, 3})
	expected = []string{"b", "d"}

	if !areStringSlicesEqual(result, expected) {
		t.Error("RebuildSliceWithIndexes function test failed.")
	}

}

func TestSplitNodeLine(t *testing.T) {

	result1, result2, result3 := SplitNodeLine("AAA = (BBB, CCC)")
	expected1 := "AAA"
	expected2 := "BBB"
	expected3 := "CCC"

	if result1 != expected1 {
		t.Errorf("SplitNodeLine function test failed. Expected: %s, Got: %s", expected1, result1)
	}
	if result2 != expected2 {
		t.Errorf("SplitNodeLine function test failed. Expected: %s, Got: %s", expected2, result2)
	}
	if result3 != expected3 {
		t.Errorf("SplitNodeLine function test failed. Expected: %s, Got: %s", expected3, result3)
	}

}

func TestLowesCommonMultipler(t *testing.T) {
	result := LowesCommonMultipler([]int{2, 3, 4})
	expected := 12

	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`)

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
	expected := 6

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
