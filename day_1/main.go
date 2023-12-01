package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetCalibrationValue(calibrationValue string) int{

	integer := 0
	twoDigitNumber := ""

	for i := 0; i < len(calibrationValue); i++{

		char := calibrationValue[i]

		if 47 < char && char < 58 { // We have a number
			
			integer = int(char) - 48
			twoDigitNumber = twoDigitNumber + strconv.Itoa(integer)

			break // Number has been found
			
		}

	}

	for i := len(calibrationValue) - 1; 0 <= i; i--{

		char := calibrationValue[i]

		if 47 < char && char < 58 { // We have a number
			
			integer = int(char) - 48
			twoDigitNumber = twoDigitNumber + strconv.Itoa(integer)

			break // Number has been found
		}

	}

	integer, err := strconv.Atoi(twoDigitNumber)

	if err != nil{
		fmt.Println(err)
		return -1
	}

	return integer
}

func AddNumbersInSlice(values []int) int{
	
	sum := 0

	for i := 0; i < len(values); i++{
		sum = sum + values[i]
	}

	return sum
}


func ReadFileContents(filename string) (string, error) {
    content, err := os.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

func main() {
	
	numbers := []int{}
	fileContent, err := ReadFileContents("puzzle_input.txt")

	if err != nil{
		fmt.Println(err)
		return
	}

	fileLines := strings.Split(fileContent, "\n")

	for _, line  := range fileLines {		
		numbers = append(numbers, GetCalibrationValue(line))
	}

	result := AddNumbersInSlice(numbers)

	fmt.Println(result)
}