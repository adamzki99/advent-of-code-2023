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

func StringOfNumbersToSingle(stringOfNumbers, seperator string) int {

	seperateValues := strings.Split(stringOfNumbers, seperator)[1:]

	concatinatedNumbers := ""

	for _, v := range seperateValues {

		if v != "" {
			concatinatedNumbers = concatinatedNumbers + v
		}

	}

	number, err := strconv.Atoi(concatinatedNumbers)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	return number
}

func ExtractRace(fileLines *[]string) Race {

	time := StringOfNumbersToSingle((*fileLines)[0], " ")
	distance := StringOfNumbersToSingle((*fileLines)[1], " ")

	return Race{time: time, distance: distance}
}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	waysOfBeatingTheRace := 0

	fileLines := strings.Split(fileContent, "\n")

	race := ExtractRace(&fileLines)

	for holdTime := 0; holdTime < race.time; holdTime++ {

		distanceTraveled := GetDistanceTraveled(holdTime, race.time)

		if distanceTraveled > race.distance {
			waysOfBeatingTheRace++
		}

	}

	return waysOfBeatingTheRace

}

func main() {

	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
