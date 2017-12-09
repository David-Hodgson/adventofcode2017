package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)

type  Program struct {
	name     string
	parent   string
	children []string
	childWeights []int
	weight   int
	totalWeight int
}

func parseStringToProgram(input string) Program {

	parts := strings.Split(input, " ")

	name := parts[0]
	weight,_ := strconv.Atoi(parts[1][1:len(parts[1])-1])

	newProgram := Program{}
	newProgram.name = name
	newProgram.weight = weight

	if len(parts) > 2 {

		subProgramList := make([]string,len(parts)-3)

		progNumber :=0
		for i := 3 ; i < len(parts) ; i++ {

			subprogramName := parts[i]
			if strings.HasSuffix(subprogramName, ",") {
				subprogramName = subprogramName[0:len(subprogramName)-1]
			}

			subProgramList[progNumber] = subprogramName
			progNumber++
		}
		newProgram.children = subProgramList
	}

	return newProgram
}

func findRoot(programStack map[string]Program) Program {

	for _,program := range programStack {

		if len(program.children) > 0 {
			for i :=0 ; i< len(program.children) ; i++ {

				childProgram := programStack[program.children[i]]
				childProgram.parent = program.name
				programStack[program.children[i]] = childProgram
			}
		}
	}

	var root Program
	for _,program := range programStack {

		if program.parent == "" {
			root = program
		}
	}

	return root
}

func DaySevenExample(){

	fmt.Println("Day Seven - Example")

	input := "pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\nktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\npadx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\njptl (61)\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\ncntj (57)"

	programInputArray := strings.Split(input, "\n")

	programStack := make(map[string]Program)

	for i := 0 ; i <len(programInputArray) ; i++ {
		program := parseStringToProgram(programInputArray[i])
		programStack[program.name] = program
	}

	fmt.Println("Root Program:", findRoot(programStack).name)
}

func DaySevenPartOne(){

	fmt.Println("Day Seven - Part One")

	input := ReadFile("day7-input.txt")
	programInputArray := strings.Split(input, "\n")

	programStack := make(map[string]Program)

	for i := 0 ; i <len(programInputArray) ; i++ {
		program := parseStringToProgram(programInputArray[i])
		programStack[program.name] = program
	}

	fmt.Println("Root Program:", findRoot(programStack).name)
}

func getTotalWeight(program Program, programStack map[string]Program) int {


	weight := program.weight

	for i:= 0; i<len(program.children); i++ {
		weight += getTotalWeight(programStack[program.children[i]], programStack)
	}

	program.totalWeight = weight
	programStack[program.name] = program

	return weight
}

func findUnbalancedPrograms(rootProgram Program, programStack map[string]Program) {

	//if we have two children either could be unbalanced
	if len(rootProgram.children) == 2 {
		fmt.Println("Program", rootProgram.name, "has two children")
	} else {

		childWeights := make(map[int]int)

		for i := 0 ; i < len(rootProgram.children); i++ {
			childProgram := programStack[rootProgram.children[i]]
			childWeights[childProgram.totalWeight]++
		}

		fmt.Println(childWeights)

		if len(childWeights) == 1 {
			fmt.Println("Program", rootProgram.name, "has incorrect amount")
		} else {
			for value, count := range childWeights {

				if count == 1 {
					for j := 0; j < len(rootProgram.children); j++ {
						childProgram := programStack[rootProgram.children[j]]
						if childProgram.totalWeight == value {
							fmt.Println("Program", childProgram.name, "is different with totalWeight", value, ".Program weight:", childProgram.weight)
							findUnbalancedPrograms(childProgram, programStack)
						}
					}
				}
			}
		}
	}
}

func DaySevenPartTwo(){

	fmt.Println("Day Seven - Part Two")

	input := ReadFile("day7-input.txt")
	programInputArray := strings.Split(input, "\n")

	programStack := make(map[string]Program)

	for i := 0 ; i <len(programInputArray) ; i++ {
		program := parseStringToProgram(programInputArray[i])
		programStack[program.name] = program
	}

	rootName := findRoot(programStack).name
	rootProgram := programStack[rootName]
	fmt.Println("Root Program:", rootName)

	fmt.Println("Root wieght:", getTotalWeight(rootProgram, programStack))
	findUnbalancedPrograms(rootProgram, programStack)
}