package main

import "fmt"

// Function with parameters
func Add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println("Hello, World!")

	// Calling the function with arguments
    result := Add(3, 7)
    
    // Printing the result
    fmt.Println("Result:", result)
}