package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Race struct {
	time     int
	distance int
}

func GetDistanceTraveled(holdTime, raceTime int) int {

	travelTime := raceTime - holdTime

	if -1 < travelTime {
		return holdTime * travelTime
	}

	return 0
}

func StringOfNumbersToSliceOfNumbers(stringOfNumbers, seperator string) []int {

	seperateValues := strings.Split(stringOfNumbers, seperator)[1:]

	returnSlice := []int{}

	for _, v := range seperateValues {

		if v != "" {
			convertedIntegers, err := strconv.Atoi(v)

			if err != nil {
				return []int{}
			}

			returnSlice = append(returnSlice, convertedIntegers)
		}

	}

	return returnSlice
}

func ExtractRaces(fileLines *[]string) []Race {

	returnSlice := []Race{}

	times := StringOfNumbersToSliceOfNumbers((*fileLines)[0], " ")
	distances := StringOfNumbersToSliceOfNumbers((*fileLines)[1], " ")

	for i := range times {
		returnSlice = append(returnSlice, Race{time: times[i], distance: distances[i]})

	}

	return returnSlice
}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	puzzleAwnser := 1
	waysOfBeatingTheRace := 0

	fileLines := strings.Split(fileContent, "\n")

	for _, race := range ExtractRaces(&fileLines) {

		for holdTime := 0; holdTime < race.time; holdTime++ {

			distanceTraveled := GetDistanceTraveled(holdTime, race.time)

			if distanceTraveled > race.distance {
				waysOfBeatingTheRace++
			}

		}

		puzzleAwnser = puzzleAwnser * waysOfBeatingTheRace
		waysOfBeatingTheRace = 0

	}

	return puzzleAwnser

}

func main() {

	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
