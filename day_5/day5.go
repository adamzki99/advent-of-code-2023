package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/adamzki99/advent-of-code-2023/packages/file"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Last() *Node {

	currentNode := ll.head

	for {

		if currentNode.next == nil {
			return currentNode
		}

		currentNode = currentNode.next
	}

}

type Mapping struct {
	lowerBound int
	upperBound int
	change     int
}

type Range struct {
	lower int
	upper int
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
			lowerBound: extractedValuesFromLine[1],
			upperBound: extractedValuesFromLine[1] + extractedValuesFromLine[2] - 1,
			change:     extractedValuesFromLine[0] - extractedValuesFromLine[1],
		}

		mappings = append(mappings, newMapping)
	}

	return mappings
}

func CreateSliceOfSeeds(fileContent *string) []int {

	fileLines := strings.Split(*fileContent, "\n")

	return StringOfNumbersToSliceOfNumbers(strings.Split(fileLines[0], ": ")[1], " ")

}

func CreateSliceOfSeedRanges(fileContent *string) []Range {

	returnSlice := []Range{}

	numbers := CreateSliceOfSeeds(fileContent)

	for i := 0; i < len(numbers); i = i + 2 {
		returnSlice = append(returnSlice, Range{lower: numbers[i], upper: numbers[i] + numbers[i+1]})
	}

	return returnSlice
}

func GetLowestValue(seedRange Range, mapStrings []string, wg *sync.WaitGroup, resultChan chan int) {

	defer wg.Done()

	lowestValue := math.MaxInt64

	for currentSeed := seedRange.lower; currentSeed < seedRange.upper; currentSeed++ {

		if currentSeed == 82 {
			fmt.Println(currentSeed)
		}

		seedCopy := currentSeed

		for _, mapString := range mapStrings {

			mappings := CreateMappings(mapString)

			for _, mapping := range mappings {

				if mapping.lowerBound <= seedCopy && seedCopy <= mapping.upperBound {
					seedCopy = seedCopy + mapping.change
				}
			}

		}

		if currentSeed == 82 {
			fmt.Println(seedCopy)
		}

		if seedCopy < lowestValue {
			lowestValue = seedCopy
		}
	}

	resultChan <- lowestValue

}

func SolvePuzzle(fileName string) int {

	var wg sync.WaitGroup

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	mapStrings := strings.Split(fileContent, "\n\n")
	mapStrings = mapStrings[1:]

	seedRanges := CreateSliceOfSeedRanges(&fileContent)

	numberOfWorkers := len(seedRanges)
	resultChan := make(chan int, numberOfWorkers)

	for _, seedRange := range seedRanges {

		wg.Add(1)
		go GetLowestValue(seedRange, mapStrings, &wg, resultChan)

	}

	// Close the result channel once all workers are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	lowestValue := math.MaxInt64

	// Get results from the channel
	for result := range resultChan {

		if result < lowestValue {
			lowestValue = result
		}
	}

	return lowestValue
}

func main() {

	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
