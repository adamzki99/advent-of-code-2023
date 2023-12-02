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

func CheckDrawValidity(red, green, blue  int) bool{
	
	if red > 12 || green > 13 || blue > 14{
		return false
	}

	return true
}

func main(){

	validGameIDSum := 0

	fileContent, err := file.ReadFileContents("puzzle_input.txt")

	if err != nil{
		fmt.Println(err)
		return
	}

	fileContentLineByLine := strings.Split(fileContent, "\n")

	gameLoop:
	for _, line := range fileContentLineByLine {
		
		gameID := GetGameID(line)

		draws := GetDraws(line)

		for _, draw := range draws {

			redCubes := GetCubes(draw, "red")
			greenCubes := GetCubes(draw, "green")
			blueCubes := GetCubes(draw, "blue")

		
			if !CheckDrawValidity(redCubes, greenCubes, blueCubes){
				continue gameLoop
			}
		}
		validGameIDSum = validGameIDSum + gameID
	}

	fmt.Println(validGameIDSum)
}