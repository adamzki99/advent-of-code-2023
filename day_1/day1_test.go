package main

import (
	"testing"
)

func TestAddNumbersInSlice(t *testing.T) {

	testSlice := []int{12, 38, 15, 77}
	result := AddNumbersInSlice(testSlice)
	expected := 142

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("AddNumbersInSlice function test failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestFindFirstNumber(t *testing.T) {

	result := FindFirstNumber("foursevennvmgqvxhvg6seven")
	expected := "4"

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("FindFirstNumber function test failed. Expected: %s, Got: %s", expected, result)
	}

	result = FindFirstNumber("4sixbrcbxq")
	expected = "4"

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("FindFirstNumber function test failed. Expected: %s, Got: %s", expected, result)
	}

	result = FindFirstNumber("kcdflseven69")
	expected = "7"

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("FindFirstNumber function test failed. Expected: %s, Got: %s", expected, result)
	}
}

func TestFindLastNumber(t *testing.T) {

	result := FindLastNumber("foursevennvmgqvxhvg6seven")
	expected := "7"

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("FindLastNumber function test failed. Expected: %s, Got: %s", expected, result)
	}

	result = FindLastNumber("4sixbrcbxq")
	expected = "6"

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("FindLastNumber function test failed. Expected: %s, Got: %s", expected, result)
	}

	result = FindLastNumber("kcdflseven69")
	expected = "9"

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("FindLastNumber function test failed. Expected: %s, Got: %s", expected, result)
	}
}
