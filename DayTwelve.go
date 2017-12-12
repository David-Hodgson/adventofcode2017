package adventofcode2017

import (
	"fmt"
	"strings"
)

func DayTwelveExample() {

	fmt.Println("Day Twelve - Example")

	input := "0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5"

	pipeList := strings.Split(input, "\n")

	programMap := make(map[string]map[string]bool)

	for i := 0; i < len(pipeList); i++ {
		parts := strings.Split(pipeList[i], "<->")

		sendingProgram := strings.Trim(parts[0], " ")
		destinationPrograms := strings.Split(strings.Trim(parts[1], " "), ", ")

		destinationMap := make(map[string]bool)

		for j:=0 ; j< len(destinationPrograms); j++ {
			destinationMap[destinationPrograms[j]] = true
		}

		programMap[sendingProgram] = destinationMap
	}

	//fmt.Println(programMap)

	//fmt.Println(programMap["0"])
	destinationSize := len(programMap["0"])

	//fmt.Println(destinationSize)
	for  {

		programsToProcess := programMap["0"]

		for destprogram,_ := range programsToProcess {

			//fmt.Println("Processing program:", destprogram)

			for otherProgram,_ := range programMap[destprogram] {
				programMap["0"][otherProgram] = true
			}
			//fmt.Println(programMap)
		}

		if len(programMap["0"]) == destinationSize {
			//fmt.Println("Would have broke")
			break;
		}

		destinationSize = len(programMap["0"])
		//break;
	}

	fmt.Println("Program 0 destinations:", len(programMap["0"]))
}

func processProgramPipes(programToStart string, programMap map[string]map[string]bool ) {
	destinationSize := len(programMap[programToStart])

	for  {

		programsToProcess := programMap[programToStart]

		for destprogram,_ := range programsToProcess {

			//fmt.Println("Processing program:", destprogram)

			for otherProgram,_ := range programMap[destprogram] {
				programMap[programToStart][otherProgram] = true
			}
			//fmt.Println(programMap)
		}

		if len(programMap[programToStart]) == destinationSize {
			//fmt.Println("Would have broke")
			break;
		}

		destinationSize = len(programMap[programToStart])
		//break;
	}
}

func buildProgramMap(pipeList []string) map[string]map[string]bool {
	programMap := make(map[string]map[string]bool)

	for i := 0; i < len(pipeList); i++ {
		parts := strings.Split(pipeList[i], "<->")

		sendingProgram := strings.Trim(parts[0], " ")
		destinationPrograms := strings.Split(strings.Trim(parts[1], " "), ", ")

		destinationMap := make(map[string]bool)

		for j:=0 ; j< len(destinationPrograms); j++ {
			destinationMap[destinationPrograms[j]] = true
		}

		programMap[sendingProgram] = destinationMap
	}

	return programMap
}

func DayTwelvePartOne() {

	fmt.Println("Day Twelve - Part One")

	input := ReadFile("day12-input.txt")
	pipeList := strings.Split(input, "\n")

	programMap := buildProgramMap(pipeList)
	processProgramPipes("0", programMap)

	fmt.Println("Program 0 destinations:", len(programMap["0"]))
}


func DayTwelvePartTwo() {

	fmt.Println("Day Twelve - Part Two")

	input := ReadFile("day12-input.txt")
	pipeList := strings.Split(input, "\n")

	programMap := buildProgramMap(pipeList)
	programGroupMaps := make(map[string]bool)

	for destprogram,_ := range programMap {

		process := true
		for pg,_ := range programGroupMaps {
			progToCheck := programMap[pg]

			if progToCheck[destprogram] {
				process = false
			}
		}

		if process {
			processProgramPipes(destprogram, programMap)
			programGroupMaps[destprogram] = true
		}

	}

	fmt.Println("Number of groups:", len(programGroupMaps))
}
