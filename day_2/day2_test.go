package main

import (
	"testing"
)

func TestGetGameID(t *testing.T){

	result := GetGameID("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	expected := 1

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetGameID function test failed. Expected: %d, Got: %d", expected, result)
    }

	result = GetGameID("Game 12: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	expected = 12

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetGameID function test failed. Expected: %d, Got: %d", expected, result)
    }

}

func TestGetCubes(t *testing.T){

	result := GetCubes("3 blue, 4 red", "red")
	expected := 4

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetRedCubes function test failed. Expected: %d, Got: %d", expected, result)
    }

	result = GetCubes("1 red, 2 green, 6 blue", "green")
	expected = 2

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetRedCubes function test failed. Expected: %d, Got: %d", expected, result)
    }

	result = GetCubes("2 green", "blue")
	expected = 0

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetRedCubes function test failed. Expected: %d, Got: %d", expected, result)
    }
}

func TestCheckDrawValidity(t *testing.T){

	result := CheckDrawValidity(12, 13, 14)
	expected := true

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("CheckDrawValidity function test failed. Expected: %t, Got: %t", expected, result)
    }

	result = CheckDrawValidity(8, 20, 13)
	expected = false

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("CheckDrawValidity function test failed. Expected: %t, Got: %t", expected, result)
    }

	result = CheckDrawValidity(13, 8, 12)
	expected = false

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("CheckDrawValidity function test failed. Expected: %t, Got: %t", expected, result)
    }

	result = CheckDrawValidity(1, 8, 42)
	expected = false

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("CheckDrawValidity function test failed. Expected: %t, Got: %t", expected, result)
    }
}

func areSlicesEqual(slice1, slice2 []string) bool {
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

func TestGetDraws(t *testing.T) {

	result := GetDraws("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	expected := []string{"3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"}

	if !areSlicesEqual(result, expected){
		t.Errorf("GetDraws function test failed. Slices are not equal")
	}

	
}