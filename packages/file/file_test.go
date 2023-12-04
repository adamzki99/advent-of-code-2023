package file

import (
	"os"
	"testing"
)

func TestReadFileContents(t *testing.T) {
	// Create a temporary test file
	content := []byte("Hello, this is a test file.")
	tmpfile, err := os.CreateTemp("", "example")
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
