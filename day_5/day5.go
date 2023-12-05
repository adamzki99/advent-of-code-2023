package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

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

func CreateListOfSeeds(fileContent *string) []int {

	fileLines := strings.Split(*fileContent, "\n")

	return StringOfNumbersToSliceOfNumbers(strings.Split(fileLines[0], ": ")[1], " ")

}

func SolvePuzzle(fileName string) int {

	fileContent, err := file.ReadFileContents(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	listOfSeeds := []LinkedList{}

	for _, seed := range CreateListOfSeeds(&fileContent) {

		newLinkedList := LinkedList{}
		newLinkedList.Append(seed)

		listOfSeeds = append(listOfSeeds, newLinkedList)
	}

	mapStrings := strings.Split(fileContent, "\n\n")
	mapStrings = mapStrings[1:]

	for _, mapString := range mapStrings {

		mappings := CreateMappings(mapString)

	seedLoop:
		for _, seed := range listOfSeeds {

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
}

func main() {

	fmt.Println(SolvePuzzle("puzzle_input.txt"))
}
