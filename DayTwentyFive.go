package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)

type transition struct {
	writeValue int
	move int
	newState string
}

func DayTwentyFiveExample() {

	fmt.Println("Day 25 - Example")

	steps := 6
	tape := make(map[int]int)
	tapePos := 0
	currentState := "A"

	stateRules := make(map[string][]transition)

	aTransition0 := transition{1,RIGHT,"B"}
	aTransition1 := transition{0,LEFT,"B"}

	aTransitions := []transition {aTransition0,aTransition1}
	stateRules["A"] = aTransitions

	bTransition0 := transition{1,LEFT,"A"}
	bTransition1 := transition{1,RIGHT,"A"}

	bTransitions := []transition {bTransition0,bTransition1}
	stateRules["B"] = bTransitions

	for i:=0;i<steps;i++ {

		fmt.Println()
		fmt.Println("Step:", i)

		fmt.Println("CurrentState:", currentState)

		fmt.Println("Tape Pos:", tapePos)
		fmt.Println("Tape Value:", tape[tapePos])

		currentTransition := stateRules[currentState][tape[tapePos]]
		fmt.Println(currentTransition)

		tape[tapePos] = currentTransition.writeValue
		if currentTransition.move == RIGHT {
			tapePos++
		} else {
			tapePos--
		}

		currentState = currentTransition.newState
		fmt.Println(tape)
	}

	fmt.Println(tape)
	fmt.Println("Done")
}

func getTransition(lines []string, offset int) transition{

	valueLine := lines[offset+1]
	moveLine := lines[offset+2]
	stateLine := lines[offset+3]

	value,_ := strconv.Atoi(valueLine[len(valueLine)-2:len(valueLine)-1])
	state := stateLine[len(stateLine)-2:len(stateLine)-1]
	move := RIGHT

	if strings.HasSuffix(moveLine, "left.") {
		move = LEFT
	}

	return transition{value,move,state}
}

func parseStateRules(input string) map[string][]transition {

	lines := strings.Split(input, "\n")

	stateRules := make(map[string][]transition)

	for i:=0 ; i< len(lines) ; i++ {
		currentLine := lines[i]

		if strings.HasPrefix(currentLine, "In state") {
			state := currentLine[len(currentLine)-2:len(currentLine)-1]
			transition0 := getTransition(lines,i+1)
			transition1 := getTransition(lines,i+5)

			stateTransitions := []transition {transition0,transition1}
			stateRules[state] = stateTransitions


		}
	}
	return stateRules
}

func DayTwentyFivePartOne() {

	fmt.Println("Day 25 - Part One")

	input := ReadFile("day25-input.txt")

	stateRules := parseStateRules(input)
	fmt.Println(stateRules)
	steps := 12173597
	tape := make(map[int]int)
	tapePos := 0
	currentState := "A"

	for i:=0;i<steps;i++ {

		currentTransition := stateRules[currentState][tape[tapePos]]

		tape[tapePos] = currentTransition.writeValue
		if currentTransition.move == RIGHT {
			tapePos++
		} else {
			tapePos--
		}

		currentState = currentTransition.newState
	}

	fmt.Println(tape)

	oneCount := 0

	for _,value := range tape {

		oneCount += value
	}

	fmt.Println("Count:", oneCount)
	fmt.Println("Done")
}
