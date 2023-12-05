package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Range struct {
	lower int
	upper int
}

type Node struct {
	seedRange Range
	children  []Node
}

type Mapping struct {
	effectRange Range
	change      int
}

func StringOfNumbersToSliceOfNumbers(stringOfNumbers, seperator string) []int {

	seperateValues := strings.Split(stringOfNumbers, seperator)

	returnSlice := []int{}

	for _, v := range seperateValues {

		convertedIntegers, err := strconv.Atoi(v)

		if err != nil {
			return []int{}
		}

		returnSlice = append(returnSlice, convertedIntegers)
	}

	return returnSlice
}

func CreateMappings(mapInput string) []Mapping {

	mappings := []Mapping{}

	mapInputLines := strings.Split(mapInput, "\n")
	mapInputLines = mapInputLines[1:]

	for _, line := range mapInputLines {
		extractedValuesFromLine := StringOfNumbersToSliceOfNumbers(line, " ")

		if len(extractedValuesFromLine) == 0 {
			return []Mapping{}
		}

		newMapping := Mapping{
			effectRange: Range{lower: extractedValuesFromLine[1], upper: (extractedValuesFromLine[1] + extractedValuesFromLine[2] - 1)},
			change:      extractedValuesFromLine[0] - extractedValuesFromLine[1],
		}

		mappings = append(mappings, newMapping)
	}

	return mappings
}

func CreateSliceOfSeedRanges(fileContent *string) []Range {

	returnList := []Range{}

	fileLines := strings.Split(*fileContent, "\n")

	numbers := StringOfNumbersToSliceOfNumbers(strings.Split(fileLines[0], ": ")[1], " ")

	for i := 0; i < len(numbers); i = i + 2 {

		returnList = append(returnList, Range{lower: numbers[i], upper: numbers[i] + numbers[i+1] - 1})

	}

	return returnList

}

// The "breakpoint" is the upperbound for the lower divide
func BreakUpRange(r Range, breakPoint, flagLowerUpper int) Range {

	if flagLowerUpper == 1 { //upper
		return Range{lower: breakPoint + 1, upper: r.upper}
	} else { // lower
		return Range{lower: r.lower, upper: breakPoint}
	}
}

func DoesEffectRangeContainSeedRangeCompletly(effectRange, seedRange Range) bool {

	if effectRange.lower <= seedRange.lower && seedRange.upper <= effectRange.upper {
		return true
	}
	return false
}

// Returns the last seed in the seedrange that is effected by effectRange
func DoesEffectRangeContianSeedRangePartially(effectRange, seedRange Range) (int, int) {

	if effectRange.lower <= seedRange.lower && effectRange.upper < seedRange.upper {
		return effectRange.upper, 1
	}
	if seedRange.lower < effectRange.lower && seedRange.upper <= effectRange.upper {
		return effectRange.lower, -1
	}

	return -1, 0
}

func PopulateSubTree(mapStrings []string, subTreeHead *Node) {

	if len(mapStrings) == 0 {
		return
	}

	mapString := mapStrings[0]

	mappings := CreateMappings(mapString)

	for _, mapping := range mappings {

		child := Node{}

		if DoesEffectRangeContainSeedRangeCompletly(mapping.effectRange, subTreeHead.seedRange) {

			child.seedRange = Range{
				lower: subTreeHead.seedRange.lower + mapping.change,
				upper: subTreeHead.seedRange.upper + mapping.change,
			}

		} else {

			breakingPoint, flag := DoesEffectRangeContianSeedRangePartially(mapping.effectRange, subTreeHead.seedRange)

			if breakingPoint != -1 {

				effectedRange := BreakUpRange(subTreeHead.seedRange, breakingPoint, flag)

				child.seedRange = Range{
					lower: effectedRange.lower + mapping.change,
					upper: effectedRange.upper + mapping.change,
				}

				PopulateSubTree(mapStrings[1:], &child)

				subTreeHead.children = append(subTreeHead.children, child)

			}

			child.seedRange = Range{
				lower: subTreeHead.seedRange.lower,
				upper: subTreeHead.seedRange.upper,
			}

		}

		PopulateSubTree(mapStrings[1:], &child)

		subTreeHead.children = append(subTreeHead.children, child)

	}

}

func GetLowestSeedInTree(n *Node) int {

	localReturn := n.seedRange.lower

	for _, child := range n.children {

		functionReturn := GetLowestSeedInTree(&child)

		if functionReturn < localReturn && len(n.children) == 0 {
			localReturn = functionReturn
		}
	}

	return localReturn
}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	seedTree := Node{children: []Node{}}

	for _, seedRange := range CreateSliceOfSeedRanges(&fileContent) {

		newNode := Node{seedRange: seedRange}
		seedTree.children = append(seedTree.children, newNode)
	}

	mapStrings := strings.Split(fileContent, "\n\n")
	mapStrings = mapStrings[1:]

	PopulateSubTree(mapStrings, &seedTree)

	lowestValue := GetLowestSeedInTree(&seedTree)

	return lowestValue

<<<<<<< HEAD
=======
			for _, mapping := range mappings {

				if mapping.lowerBound <= seed.Last().data && seed.Last().data <= mapping.upperBound {
					seed.Append(seed.Last().data + mapping.change)
					continue seedLoop
				}

			}

			seed.Append(seed.Last().data)

		}
	}

	//for _, seed := range listOfSeeds {

	//seed.Display()

	//}

	// Now we can get the lowest location

	lowestLocation := math.MaxUint32

	for _, seed := range listOfSeeds {

		//fmt.Println(seed.Last().data)

		if seed.Last().data < lowestLocation {
			lowestLocation = seed.Last().data
		}

	}

	return lowestLocation
>>>>>>> parent of 63b32fd (Part 1 complete)
}

func main() {

	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
