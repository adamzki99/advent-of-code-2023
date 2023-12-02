package main

import (
	"testing"
	"os"
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

func TestRunProgram(t *testing.T){

	content := []byte("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")
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

	result := RunProgram(tmpfile.Name())
	expected := 2286

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("RunProgram function test failed. Expected: %d, Got: %d", expected, result)
    }

}