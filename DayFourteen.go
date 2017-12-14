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

func buildGrid(inputKey string) [][]bool {

	grid := make([][]bool,128)

	for i := 0; i<128 ; i++ {

		rowKey := inputKey + "-" + strconv.Itoa(i)

		grid[i] = make([]bool, 128)

		knotHash := fmt.Sprintf(KnotHash(rowKey))

		col := 0
		for j := 0 ; j < len(knotHash); j++ {
			charString := fmt.Sprintf("%04b",asciiToHex(string(knotHash[j])))

			for k := 0 ; k <len(charString); k++ {
				if string(charString[k]) == "1" {
					grid[i][col] = true
				}
				col++
			}
		}

	}

	return grid
}

func printGrid(grid[][]bool) {
	for i := 0 ; i<len(grid); i++ {

		for j := 0 ; j < len(grid[i]); j++ {
			if grid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func isNewGroup(i,j int, grid [][]bool,groupedPoints map[string]bool) bool {

	if !grid[i][j] {
		return false
	}

	groupKey := getGroupKey(i,j)

	if groupedPoints[groupKey] {
		return false
	}

	return true
}

func getGroupKey(i,j int) string {
	return strconv.Itoa(i) + "," + strconv.Itoa(j)
}

func findGroupMembers(i,j int, grid [][]bool, groupedPoints map[string]bool) {

	//above
	if i > 0 {
		if isNewGroup(i-1,j,grid,groupedPoints) {
			groupedPoints[getGroupKey(i-1,j)] = true
			findGroupMembers(i-1,j,grid,groupedPoints)
		}
	}
	//below
	if i < 127 {
		if isNewGroup(i+1,j,grid,groupedPoints) {
			groupedPoints[getGroupKey(i+1,j)] = true
			findGroupMembers(i+1,j,grid,groupedPoints)
		}
	}
	//left
	if j > 0 {
		if isNewGroup(i,j-1,grid,groupedPoints) {
			groupedPoints[getGroupKey(i,j-1)] = true
			findGroupMembers(i,j-1,grid,groupedPoints)
		}
	}

	//right
	if j < 127 {
		if isNewGroup(i,j+1,grid,groupedPoints) {
			groupedPoints[getGroupKey(i,j+1)] = true
			findGroupMembers(i,j+1,grid,groupedPoints)
		}
	}
}

func findGroups(grid[][]bool) int {
	groupedPoints := make(map[string]bool)

	groupCount := 0
	for i := 0 ; i<len(grid); i++ {

		for j := 0 ; j < len(grid[i]); j++ {

			//Test grid i,j
			//fmt.Println("Testing",i,",",j)
			isNewGroup := isNewGroup(i,j,grid, groupedPoints)

			if isNewGroup {
				findGroupMembers(i,j,grid,groupedPoints)
				groupCount++
			}
			//fmt.Println("\tNew group:", isNewGroup)
		}
	}

	return groupCount
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

func DayFourteenPartTwo() {

	fmt.Println("Day 14 - Part Two")

	exampleInput := "amgozmfv"

	grid := buildGrid(exampleInput)

	fmt.Println("Group count:", findGroups(grid))
}
