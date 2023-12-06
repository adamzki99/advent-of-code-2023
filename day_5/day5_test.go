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

func areMappingsEqual(m1, m2 Mapping) bool {

	if m1.lowerBound != m2.lowerBound {
		return false
	}

	if m1.upperBound != m2.upperBound {
		return false
	}

	if m1.change != m2.change {
		return false
	}

	return true

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

func TestCreateMapping(t *testing.T) {

	input := `seed-to-soil map:
50 98 2
52 50 48`

	result := CreateMappings(input)
	expected := []Mapping{
		{

			lowerBound: 98,
			upperBound: 98 + 2 - 1,
			change:     50 - 98,
		},
		{

			lowerBound: 50,
			upperBound: 50 + 48 - 1,
			change:     52 - 50,
		},
	}

	if len(result) != len(expected) {
		t.Errorf("CreateMappings function test failed.")
	}

	for i := range result {

		if !areMappingsEqual(result[i], expected[i]) {
			t.Errorf("CreateMappings function test failed.")
		}

	}

	input = `soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15`

	result = CreateMappings(input)
	expected = []Mapping{
		{
			lowerBound: 15,
			upperBound: 15 + 37 - 1,
			change:     0 - 15,
		},
		{
			lowerBound: 52,
			upperBound: 52 + 2 - 1,
			change:     37 - 52,
		},
		{
			lowerBound: 0,
			upperBound: 0 + 15 - 1,
			change:     39 - 0,
		},
	}

	for i := range result {

		if !areMappingsEqual(result[i], expected[i]) {
			t.Errorf("CreateMappings function test failed.")
		}

	}
}

func TestCreateSliceOfSeedRanges(t *testing.T) {

	content := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	result := CreateSliceOfSeedRanges(&content)
	expected := []Range{
		{
			lower: 79,
			upper: 92,
		},
		{
			lower: 55,
			upper: 67,
		},
	}

	index := 0
	if result[index].lower != expected[index].lower || result[index].upper != expected[index].upper {
		t.Error("CreateSliceOfSeedRanges function test failed.")
	}
	index = 1
	if result[index].lower != expected[index].lower || result[index].upper != expected[index].upper {
		t.Error("CreateSliceOfSeedRanges function test failed.")
	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
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
	expected := 56

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
