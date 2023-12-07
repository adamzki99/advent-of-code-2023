package main

import (
	"os"
	"testing"
)

func TestLabelConversion(t *testing.T) {

	result := LabelConversion("H")
	expected := -1
	if result != expected {
		t.Errorf("LabelConversion function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = LabelConversion("A")
	expected = 41
	if result != expected {
		t.Errorf("LabelConversion function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = LabelConversion("7")
	expected = 13
	if result != expected {
		t.Errorf("LabelConversion function test failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestExtractOccurrences(t *testing.T) {

	result := ExtractOccurrences([]string{"A", "A", "3", "3"})
	expected := make(map[int]int)
	expected[13] = 2
	expected[2] = 2

	if result[13] != expected[13] || result[2] != expected[2] {
		t.Errorf("TestCharFactory function test failed.")
	}

}

func TestHandTypeExtraction(t *testing.T) {

	result := HandTypeExtraction("AAAAA")
	expected := 71
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = HandTypeExtraction("AA8AA")
	expected = 67
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = HandTypeExtraction("23332")
	expected = 61
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = HandTypeExtraction("23232")
	expected = 61
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = HandTypeExtraction("TTT98")
	expected = 59
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = HandTypeExtraction("23432")
	expected = 53
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = HandTypeExtraction("KK677")
	expected = 53
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = HandTypeExtraction("A23A4")
	expected = 47
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}
	result = HandTypeExtraction("23456")
	expected = 43
	if result != expected {
		t.Errorf("HandTypeExtraction function test failed. Expected: %d, Got: %d", expected, result)
	}

}

func TestCompareHands(t *testing.T) {

	result := CompareHands(
		Hand{
			hType:   59,
			hString: "QQQJA",
		},
		Hand{
			hType:   59,
			hString: "T55J5",
		},
	)
	expected := true
	if result != expected {
		t.Errorf("CompareHands function test failed. Expected: %t, Got: %t", expected, result)
	}

	result = CompareHands(
		Hand{
			hType:   53,
			hString: "KK677",
		},
		Hand{
			hType:   59,
			hString: "T55J5",
		},
	)
	expected = false
	if result != expected {
		t.Errorf("CompareHands function test failed. Expected: %t, Got: %t", expected, result)
	}

	result = CompareHands(
		Hand{
			hType:   53,
			hString: "KK677",
		},
		Hand{
			hType:   53,
			hString: "KTJJT",
		},
	)
	expected = true
	if result != expected {
		t.Errorf("CompareHands function test failed. Expected: %t, Got: %t", expected, result)
	}

	result = CompareHands(
		Hand{
			hType:   53,
			hString: "KTJJT",
		},
		Hand{
			hType:   53,
			hString: "KK677",
		},
	)
	expected = false
	if result != expected {
		t.Errorf("CompareHands function test failed. Expected: %t, Got: %t", expected, result)
	}

}

func TestSolvePuzzle(t *testing.T) {

	content := []byte(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
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
	expected := 6440

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("SolvePuzzle function test failed. Expected: %d, Got: %d", expected, result)
	}

}
