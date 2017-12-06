package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

func FindMaxValue(values []int) int {
	currentMaxValue := -1
	currentMaxPos := 0

	for i := 0 ; i< len(values) ; i++ {

		if values[i] > currentMaxValue  {
			currentMaxValue = values[i]
			currentMaxPos = i
		}
	}

	return currentMaxPos
}

func MoveValues(values []int, posToMove int) {


	blocksToMove := values[posToMove]
	//fmt.Println("Moving from pos:", posToMove, "Blocks Count:", blocksToMove)
	values[posToMove] = 0

	currentPos := posToMove
	for i :=0 ; i< blocksToMove; i++ {
		//fmt.Println("Block Count:", blocksToMove)
		currentPos++
		if currentPos > len(values) - 1 {
			currentPos = 0
		}

		values[currentPos]++
		//blocksToMove--
	}

}

func ProcessValues(values []int) int {

	//Find highest box
	currentMaxPos := FindMaxValue(values)

	//fmt.Println("Max Pos:",currentMaxPos)
	previousStates := make(map[string]bool)

	previousStates[valueArryToString(values)] = true

	step := 0
	for ; ; {

		MoveValues(values, currentMaxPos)
		step++

		if previousStates[valueArryToString(values)] == true {
			break;
		} else {
			previousStates[valueArryToString(values)] = true
			currentMaxPos = FindMaxValue(values)
		}

	}

	return step
}

func valueArryToString(values []int) string {

	var output string

	for i := 0 ; i < len(values) ; i++ {

		output += string(values[i])
		output += " "
	}

	return output
}

func DaySixExample() {


	fmt.Println("Day Six - Example")
	values := []int {0, 2, 7,  0}


	fmt.Println(values)

	fmt.Println(ProcessValues(values))
	fmt.Println(ProcessValues(values))

}

func DaySixPartOne() {

	fmt.Println("Day Six - Part One")

	input := "4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3"

	valuesAsStrings := strings.Split(input, "\t")
	values := make([]int, len(valuesAsStrings))

	for i :=0 ; i<len(valuesAsStrings); i++ {

		values[i],_ = strconv.Atoi(valuesAsStrings[i])

	}

	fmt.Println(values)

	fmt.Println("Part One:", ProcessValues(values))
}

func DaySixPartTwo() {

	fmt.Println("Day Six - Part Two")

	input := "4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3"

	valuesAsStrings := strings.Split(input, "\t")
	values := make([]int, len(valuesAsStrings))

	for i :=0 ; i<len(valuesAsStrings); i++ {

		values[i],_ = strconv.Atoi(valuesAsStrings[i])

	}

	fmt.Println(values)

	fmt.Println("Part One:", ProcessValues(values))
	fmt.Println("Part Two:",ProcessValues(values))

}

