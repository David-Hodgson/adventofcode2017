package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)

var (
	absMax int
)

func doCheck(value1, value2 int, opCode string) bool  {

	switch opCode {
	case "==":
		return value1 == value2
	case "!=":
		return value1 != value2
	case ">=":
		return value1 >= value2
	case "<=":
		return value1 <= value2
	case ">":
		return value1 > value2
	case "<":
		return value1 < value2
	}

	return false
}

func processInstructions(instructions []string) map[string] int {
	registers := make(map[string]int)

	for i :=0 ; i<len(instructions); i++ {
		parts := strings.Split(instructions[i], " ");

		targetRegister := parts[0]
		opCode := parts[1]
		inc,_ := strconv.Atoi(parts[2])

		if opCode == "dec" {
			inc = inc * -1
		}
		checkRegister := parts[4]
		checkRegisterValue := registers[checkRegister]
		checkOp := parts[5]
		checkValue,_ := strconv.Atoi(parts[6])

		//fmt.Println(opCode)
		//fmt.Println(inc)
		//fmt.Println(checkRegister)
		//fmt.Println(checkRegisterValue)
		//fmt.Println("Check Op: ", checkOp)
		//fmt.Println(checkValue)
		if doCheck(checkRegisterValue, checkValue, checkOp) {
			registers[targetRegister] = registers[targetRegister] + inc

			if (registers[targetRegister] > absMax) {
				absMax = registers[targetRegister]
			}
		}
	}

	//fmt.Println(registers)
	return registers
}

func DayEightExample() {

	fmt.Println("Day Eight Example")

	input := "b inc 5 if a > 1\na inc 1 if b < 5\nc dec -10 if a >= 1\nc inc -20 if c == 10"
	instructions := strings.Split(input, "\n")
	fmt.Println(instructions)
	registers := processInstructions(instructions)
	fmt.Println(registers)
}

func DayEightPartOne() {
	fmt.Println("Day Eight Part One")

	input := ReadFile("day8-input.txt")
	instructions := strings.Split(input, "\n")
	registers := processInstructions(instructions)

	maxValue := 0

	for _, value := range registers {

		if value > maxValue {
			maxValue = value
		}
	}

	fmt.Print("Max Value:", maxValue)
}

func DayEightPartTwo() {
	fmt.Println("Day Eight Part Two")

	input := ReadFile("day8-input.txt")
	instructions := strings.Split(input, "\n")
	registers := processInstructions(instructions)

	maxValue := 0

	for _, value := range registers {

		if value > maxValue {
			maxValue = value
		}
	}

	fmt.Println("Max Value:", maxValue)
	fmt.Println("Max Value During process", absMax)
}