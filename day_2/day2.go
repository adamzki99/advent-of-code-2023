package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

func GetGameID(gameLine string) int{

	gameTitleID := strings.Split(gameLine, ":")[0]

	gameID := strings.Split(gameTitleID, " ")[1]

	integerID, _ := strconv.Atoi(gameID)

	return integerID

}

func GetDraws(gameLine string) []string{

	allGames := strings.Split(gameLine, ": ")[1]

	draws := strings.Split(allGames, "; ")

	return draws
}

func GetCubes(draw, color string) int{

	totalNumberOfCubes := 0

	cubes := strings.Split(draw, ", ")
		
	for _, cube := range cubes {
		
		if strings.Contains(cube, color){
			nrOfCubesInInstance, _:= strconv.Atoi(strings.Split(cube, " ")[0])

			totalNumberOfCubes = totalNumberOfCubes + nrOfCubesInInstance
		}
		
	}
	
	return totalNumberOfCubes
}

func RunProgram(fileName string) int {

	awnserSum := 0

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil{
		fmt.Println(err)
		return -1
	}

	fileContentLineByLine := strings.Split(fileContent, "\n")

	for _, line := range fileContentLineByLine {
		
		largestAmountOfRedCubes := 0
		largestAmountOfGreenCubes := 0
		largestAmountOfBlueCubes := 0

		draws := GetDraws(line)

		for _, draw := range draws {

			redCubes := GetCubes(draw, "red")
			greenCubes := GetCubes(draw, "green")
			blueCubes := GetCubes(draw, "blue")
			

			if largestAmountOfRedCubes < redCubes{
				largestAmountOfRedCubes = redCubes
			}
			if largestAmountOfGreenCubes < greenCubes{
				largestAmountOfGreenCubes = greenCubes
			}
			if largestAmountOfBlueCubes < blueCubes{
				largestAmountOfBlueCubes = blueCubes
			}
			
		}

		awnserSum = awnserSum + (largestAmountOfRedCubes * largestAmountOfGreenCubes * largestAmountOfBlueCubes)
	}

	return awnserSum
}

func main(){

	fmt.Println(RunProgram("puzzle_input.txt"))
}