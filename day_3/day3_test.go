package main

import (
	"os"
	"testing"
)

func areSlicesEqual(slice1, slice2 []Point) bool {
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

func areObjectSlicesEqual(slice1, slice2 []Serial) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if !areSerialObjectEqual(v, slice2[i]) {
			return false
		}
	}

	return true
}

func areSerialObjectEqual(s1, s2 Serial) bool {

	if s1.number != s2.number {
		return false
	}

	if len(s1.didgitPositions) != len(s2.didgitPositions) {
		return false
	}

	for index, position := range s1.didgitPositions {
		if position != s2.didgitPositions[index] {
			if position.X != s2.didgitPositions[index].X || position.Y != s2.didgitPositions[index].Y {
				return false
			}
		}
	}

	return true

}

func TestSymbolOnLine(t *testing.T) {

	result := SymbolsOnLine("467..114..", 0)
	expected := []Point{}

	// Check if the result matches the expected value
	if !areSlicesEqual(result, expected) {
		t.Errorf("SymbolsOnLine function test failed.")
	}
	////0123456789
	result = SymbolsOnLine("...*......", 1)
	expected = []Point{{X: 1, Y: 3}}

	// Check if the result matches the expected value
	if !areSlicesEqual(result, expected) {
		t.Errorf("SymbolsOnLine function test failed.")
	}

	result = SymbolsOnLine("617*......", 2)
	expected = []Point{{X: 2, Y: 3}}

	// Check if the result matches the expected value
	if !areSlicesEqual(result, expected) {
		t.Errorf("SymbolsOnLine function test failed.")
	}

	result = SymbolsOnLine("...$.*....", 3)
	expected = []Point{{X: 3, Y: 3}, {X: 3, Y: 5}}

	// Check if the result matches the expected value
	if !areSlicesEqual(result, expected) {
		t.Errorf("SymbolsOnLine function test failed.")
	}

	result = SymbolsOnLine("...$.*..$.", 4)
	expected = []Point{{X: 4, Y: 3}, {X: 4, Y: 5}, {X: 4, Y: 8}}

	// Check if the result matches the expected value
	if !areSlicesEqual(result, expected) {
		t.Errorf("SymbolsOnLine function test failed.")
	}
}

func TestGetNumbersAndPositons(t *testing.T) {

	//0123456789
	result := GetNumbersAndPositons("467..114..", 0)

	expectedDidgitPositions1 := []Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}}
	expectedDidgitPositions2 := []Point{{X: 0, Y: 5}, {X: 0, Y: 6}, {X: 0, Y: 7}}
	expected := []Serial{{number: 467, didgitPositions: expectedDidgitPositions1}, {number: 114, didgitPositions: expectedDidgitPositions2}}

	// Check if the result matches the expected value
	if !areObjectSlicesEqual(result, expected) {
		t.Errorf("GetNumbersAndPositons function test 1 failed.")
	}

	//0123456789
	result = GetNumbersAndPositons("..35..633.", 3)

	expectedDidgitPositions1 = []Point{{X: 3, Y: 2}, {X: 3, Y: 3}}
	expectedDidgitPositions2 = []Point{{X: 3, Y: 6}, {X: 3, Y: 7}, {X: 3, Y: 8}}
	expected = []Serial{{number: 35, didgitPositions: expectedDidgitPositions1}, {number: 633, didgitPositions: expectedDidgitPositions2}}

	// Check if the result matches the expected value
	if !areObjectSlicesEqual(result, expected) {
		t.Errorf("GetNumbersAndPositons function test 2 failed.")
	}

}

func TestConvertToNumber(t *testing.T) {

	result := ConvertToNumber([]rune{'4', '6', '7'})
	expected := 467

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("ConvertToNumber function test failed. Expected: %d, Got: %d", expected, result)
	}

	result = ConvertToNumber([]rune{'4'})
	expected = 4

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("ConvertToNumber function test failed. Expected: %d, Got: %d", expected, result)
	}

}

func arePointSlicesEqual(p1, p2 []Point) bool {

	if len(p1) != len(p2) {
		return false
	}

	for i := 0; i < len(p1); i++ {

		if p1[i].X != p2[i].X || p1[i].Y != p2[i].Y {
			return false
		}

	}

	return true

}

func TestGetSurroundingPoints(t *testing.T) {

	result := GetSurroundingPoints(Point{X: 0, Y: 0})
	expected := []Point{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}}

	// Check if the result matches the expected value
	if !arePointSlicesEqual(result, expected) {
		t.Errorf("GetSurroundingPoints function test failed.")
	}

	result = GetSurroundingPoints(Point{X: 3, Y: 4})
	expected = []Point{{X: 2, Y: 4}, {X: 2, Y: 5}, {X: 3, Y: 5}, {X: 4, Y: 5}, {X: 4, Y: 4}, {X: 4, Y: 3}, {X: 3, Y: 3}, {X: 2, Y: 3}}

	// Check if the result matches the expected value
	if !arePointSlicesEqual(result, expected) {
		t.Errorf("GetSurroundingPoints function test failed.")
	}

}

func TestIsSerial(t *testing.T) {

	symbols := []Point{{X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}, {X: 3, Y: 5}, {X: 4, Y: 3}, {X: 4, Y: 5}, {X: 4, Y: 8}}

	serial := Serial{number: 467, didgitPositions: []Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}}}
	result := IsSerial(serial, symbols)
	expected := true

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("IsSerial function test failed. Expected: %t, Got: %t", expected, result)
	}

	serial = Serial{number: 114, didgitPositions: []Point{{X: 0, Y: 5}, {X: 0, Y: 6}, {X: 0, Y: 7}}}
	result = IsSerial(serial, symbols)
	expected = false

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("IsSerial function test failed. Expected: %t, Got: %t", expected, result)
	}

}

func TestRunProgram(t *testing.T) {

	content := []byte("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")
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
	expected := 4361

	// Check if the result matches the expected value
	if result != expected {
		t.Errorf("RunProgram function test failed. Expected: %d, Got: %d", expected, result)
	}

}
