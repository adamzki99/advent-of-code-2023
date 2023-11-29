
package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(3, 7)
    expected := 10

    // Check if the result matches the expected value
    if result != expected {
        t.Errorf("Add function test failed. Expected: %d, Got: %d", expected, result)
    }
}