package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

func asciiToHex(input string) int {
	switch input {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "a":
		return 10
	case "b":
		return 11
	case "c":
		return 12
	case "d":
		return 13
	case "e":
		return 14
	case "f":
		return 15
	default:
		return 0
	}
}

func getUsedCount(inputKey string) int {

	usedCount := 0

	for i := 0; i<128 ; i++ {

		rowKey := inputKey + "-" + strconv.Itoa(i)

		knotHash := fmt.Sprintf(KnotHash(rowKey))

		for j := 0 ; j < len(knotHash); j++ {
			charString := fmt.Sprintf("%04b",asciiToHex(string(knotHash[j])))
			strippedString := strings.Replace(charString, "0", "",-1 )
			usedCount += len ( strippedString)
		}

	}

	return usedCount
}

func DayFourteenExample() {

	fmt.Println("Day 14 - Example")

	exampleInput := "flqrgnkx"

	usedCount := getUsedCount(exampleInput)
	fmt.Println("Used Count:", usedCount)
}

func DayFourteenPartOne() {

	fmt.Println("Day 14 - Part One")

	exampleInput := "amgozmfv"

	usedCount := getUsedCount(exampleInput)
	fmt.Println("Used Count:", usedCount)
}

