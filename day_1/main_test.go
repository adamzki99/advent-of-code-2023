package main

import (
    "testing"
	"io/ioutil"
    "os"
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

func TestReadFileContents(t *testing.T) {
    // Create a temporary test file
    content := []byte("Hello, this is a test file.")
    tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        t.Fatal(err)
    }

	// Remove temporary test file with the end of the test
	defer os.Remove(tmpfile.Name())

	// Check temporary file
    if _, err := tmpfile.Write(content); err != nil {
        t.Fatal(err)
    }
    if err := tmpfile.Close(); err != nil {
        t.Fatal(err)
    }

    // Test reading the file
    result, err := ReadFileContents(tmpfile.Name())
    if err != nil {
        t.Fatalf("Error reading file: %v", err)
    }

    // Compare the expected and actual content
    if result != string(content) {
        t.Errorf("Expected %q, got %q", string(content), result)
    }
}


func TestFindFirstNumber(t *testing.T){

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

func TestFindLastNumber(t *testing.T){

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
