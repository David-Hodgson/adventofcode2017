package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)

func doCommand(command string, programList []string) []string {

	switch command[0:1] {

	case "s":
		spinLength,_ := strconv.Atoi(command[1:])
		updatedList := make([]string, len(programList))
		offset:= len(programList) - spinLength
		for i := 0 ; i < spinLength ; i++ {
			updatedList[i] = programList[i+offset]
		}

		for j := spinLength ; j < len(programList) ; j++ {
			updatedList[j] = programList[j-spinLength]
		}

		return updatedList
		break
	case "x":
		exchagePair := strings.Split(command[1:], "/")
		p1,_ := strconv.Atoi(exchagePair[0])
		p2,_ := strconv.Atoi(exchagePair[1])
		tmp := programList[p1]
		programList[p1] = programList[p2]
		programList[p2] = tmp
		break
	case "p":
		partnerPrograms := strings.Split(command[1:], "/")

		program1Name := partnerPrograms[0]
		program2Name := partnerPrograms[1]

		p1 := 0
		p2 := 0

		for k := 0 ; k < len(programList) ; k++ {
			if programList[k] == program1Name {
				p1 = k
			}

			if programList[k] == program2Name {
				p2 = k
			}
		}

		tmp := programList[p1]
		programList[p1] = programList[p2]
		programList[p2] = tmp
		break
	}

	return programList
}

func populateProgramList(size int) []string {
	start := 'a'

	programList := make([]string, size)

	for i := 0; i < size ; i++ {

		programList[i] = string(int(start) + i)
	}

	return programList
}

func DaySixteenExample() {

	fmt.Println("Day 16 - Example")

	input := "s1\nx3/4\npe/b"

	commandList := strings.Split(input, "\n")
	fmt.Println(commandList)

	programList := populateProgramList(5)

	fmt.Println(programList)
	for j := 0; j < len(commandList) ; j++ {

		programList = doCommand(commandList[j], programList)
		fmt.Println(programList)
	}
}

func DaySixteenPartOne() {

	fmt.Println("Day 16 - Part One")

	input := ReadFile("day16-input.txt")

	commandList := strings.Split(input, ",")
	fmt.Println(commandList)

	programList := populateProgramList(16)

	fmt.Println(programList)
	for j := 0; j < len(commandList) ; j++ {

		programList = doCommand(commandList[j], programList)
	}

	for k :=0 ; k <len(programList); k++ {
		fmt.Print(programList[k])
	}
	fmt.Println()
}

func IntArrayEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func DaySixteenPartTwo() {

	fmt.Println("Day 16 - Part Two")

	input := ReadFile("day16-input.txt")

	commandList := strings.Split(input, ",")
	fmt.Println(commandList)

	programList := populateProgramList(16)

	originalProgramList := populateProgramList(16);


	for i:= 0 ; i< 1000000000 ; i++ {

		if i %1000 == 0 {
			fmt.Println("Loop", i)
		}

		if i != 0 && IntArrayEquals(programList, originalProgramList) {
			fmt.Println("Array is back to how it started at loop",i)
			extraSteps := 1000000000 % 60

			i = 1000000000 - extraSteps
		}

		for j := 0; j < len(commandList); j++ {

			programList = doCommand(commandList[j], programList)
		}

	}
	for k :=0 ; k <len(programList); k++ {
		fmt.Print(programList[k])
	}
	fmt.Println()
}
