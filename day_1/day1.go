package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

func AddNumbersInSlice(values []int) int {

	sum := 0

	for i := 0; i < len(values); i++ {
		sum = sum + values[i]
	}

	return sum
}

func FindFirstNumber(calibrationValues string) string {

	numberIndex := -1
	numberValue := ""

	for i, r := range calibrationValues {

		if unicode.IsNumber(r) {
			numberIndex = i
			numberValue = string(r)
			break
		}
	}

	if strings.Contains(calibrationValues, "one") && strings.Index(calibrationValues, "one") < numberIndex {
		numberValue = "1"
		numberIndex = strings.Index(calibrationValues, "one")
	}

	if strings.Contains(calibrationValues, "two") && strings.Index(calibrationValues, "two") < numberIndex {
		numberValue = "2"
		numberIndex = strings.Index(calibrationValues, "two")
	}

	if strings.Contains(calibrationValues, "three") && strings.Index(calibrationValues, "three") < numberIndex {
		numberValue = "3"
		numberIndex = strings.Index(calibrationValues, "three")
	}

	if strings.Contains(calibrationValues, "four") && strings.Index(calibrationValues, "four") < numberIndex {
		numberValue = "4"
		numberIndex = strings.Index(calibrationValues, "four")
	}

	if strings.Contains(calibrationValues, "five") && strings.Index(calibrationValues, "five") < numberIndex {
		numberValue = "5"
		numberIndex = strings.Index(calibrationValues, "five")
	}

	if strings.Contains(calibrationValues, "six") && strings.Index(calibrationValues, "six") < numberIndex {
		numberValue = "6"
		numberIndex = strings.Index(calibrationValues, "six")
	}

	if strings.Contains(calibrationValues, "seven") && strings.Index(calibrationValues, "seven") < numberIndex {
		numberValue = "7"
		numberIndex = strings.Index(calibrationValues, "seven")
	}
	if strings.Contains(calibrationValues, "eight") && strings.Index(calibrationValues, "eight") < numberIndex {
		numberValue = "8"
		numberIndex = strings.Index(calibrationValues, "eight")
	}

	if strings.Contains(calibrationValues, "nine") && strings.Index(calibrationValues, "nine") < numberIndex {
		numberValue = "9"
	}

	return numberValue
}

func FindLastNumber(calibrationValues string) string {

	numberIndex := -1
	numberValue := ""

	for i, r := range calibrationValues {

		if unicode.IsNumber(r) {
			numberIndex = i
			numberValue = string(r)
		}
	}

	if strings.Contains(calibrationValues, "one") && strings.LastIndex(calibrationValues, "one") > numberIndex {
		numberValue = "1"
		numberIndex = strings.LastIndex(calibrationValues, "one")
	}

	if strings.Contains(calibrationValues, "two") && strings.LastIndex(calibrationValues, "two") > numberIndex {
		numberValue = "2"
		numberIndex = strings.LastIndex(calibrationValues, "two")
	}

	if strings.Contains(calibrationValues, "three") && strings.LastIndex(calibrationValues, "three") > numberIndex {
		numberValue = "3"
		numberIndex = strings.LastIndex(calibrationValues, "three")
	}

	if strings.Contains(calibrationValues, "four") && strings.LastIndex(calibrationValues, "four") > numberIndex {
		numberValue = "4"
		numberIndex = strings.LastIndex(calibrationValues, "four")
	}

	if strings.Contains(calibrationValues, "five") && strings.LastIndex(calibrationValues, "five") > numberIndex {
		numberValue = "5"
		numberIndex = strings.LastIndex(calibrationValues, "five")
	}

	if strings.Contains(calibrationValues, "six") && strings.LastIndex(calibrationValues, "six") > numberIndex {
		numberValue = "6"
		numberIndex = strings.LastIndex(calibrationValues, "six")
	}

	if strings.Contains(calibrationValues, "seven") && strings.LastIndex(calibrationValues, "seven") > numberIndex {
		numberValue = "7"
		numberIndex = strings.LastIndex(calibrationValues, "seven")
	}
	if strings.Contains(calibrationValues, "eight") && strings.LastIndex(calibrationValues, "eight") > numberIndex {
		numberValue = "8"
		numberIndex = strings.LastIndex(calibrationValues, "eight")
	}

	if strings.Contains(calibrationValues, "nine") && strings.LastIndex(calibrationValues, "nine") > numberIndex {
		numberValue = "9"
	}

	return numberValue
}

func main() {

	numbers := []int{}
	fileContent, err := file.ReadFileContents("puzzle_input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileLines := strings.Split(fileContent, "\n")

	for _, line := range fileLines {

		firstNumber := FindFirstNumber(line)
		lastNumber := FindLastNumber(line)

		number, _ := strconv.Atoi(firstNumber + lastNumber)

		numbers = append(numbers, number)

	}

	result := AddNumbersInSlice(numbers)

	fmt.Println(result)
}
