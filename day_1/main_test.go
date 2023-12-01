package main

import (
    "testing"
	"io/ioutil"
    "os"
)

func TestGetCalibrationValue(t *testing.T){

	result := GetCalibrationValue("1abc2")
	expected := 12

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetCalibrationValue function test 1 failed. Expected: %d, Got: %d", expected, result)
    }

	result = GetCalibrationValue("pqr3stu8vwx")
	expected = 38

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetCalibrationValue function test 2 failed. Expected: %d, Got: %d", expected, result)
    }

	result = GetCalibrationValue("a1b2c3d4e5f")
	expected = 15

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetCalibrationValue function test 3 failed. Expected: %d, Got: %d", expected, result)
    }

	result = GetCalibrationValue("treb7uchet")
	expected = 77

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetCalibrationValue function test 4 failed. Expected: %d, Got: %d", expected, result)
	}

	result = GetCalibrationValue("5lvlhsjkxssfour")
	expected = 55

	// Check if the result matches the expected value
    if result != expected {
        t.Errorf("GetCalibrationValue function test 4 failed. Expected: %d, Got: %d", expected, result)
    }

}

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