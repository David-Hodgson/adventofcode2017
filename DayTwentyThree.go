package adventofcode2017

import (
	"fmt"
	"strings"
)

func handleInstructions(instructionList []string, registers map[string]int64) {

	mulCount := 0
	loopCount :=0 
	hLoopCount := 0

	for i :=0 ; i< len(instructionList) ; {
		//fmt.Println("Executing:", instructionList[i], " at pos:", i)
		//fmt.Println(registers)
		loopCount++
		if  instructionList[i] == "sub h -1" {
			hLoopCount++
		}
		instructionParts := strings.Split(instructionList[i], " ")

		opCode := instructionParts[0]

		target := instructionParts[1]
		value := instructionParts[2]
		inc :=1
		switch(opCode) {
			case "set":
				registers[target] = getRegisterValue(value,registers)
				break
			case "sub":
				registers[target] = registers[target] - getRegisterValue(value,registers)
				break
			case "mul":
				mulCount++
				registers[target] = registers[target] * getRegisterValue(value,registers)
				break
			case "jnz":
				if getRegisterValue(target,registers) != 0 {
					inc = int(getRegisterValue(value,registers))
				}
				break
		}

		i += inc
	}
	fmt.Println("Mul count:", mulCount)
	fmt.Println("Loop count:", loopCount)
	fmt.Println("h loop count:" , hLoopCount)
}

func DayTwentyThreeExample() {

	fmt.Println("Day 23 - Example")

	input := "set a 1\nsub b 2"
	inputList := strings.Split(input, "\n")

	fmt.Println(inputList)
	registers := make(map[string]int64)

	fmt.Println(registers)

	handleInstructions(inputList, registers)
	fmt.Println(registers)
}

func DayTwentyThreePartOne() {

	fmt.Println("Day 23 - Part 1")

	input := ReadFile("day23-input.txt")
	inputList := strings.Split(input, "\n")

	fmt.Println(inputList)
	registers := make(map[string]int64)

	fmt.Println(registers)

	handleInstructions(inputList, registers)
	fmt.Println(registers)
}

func DayTwentyThreePartTwo() {

	fmt.Println("Day 23 - Part 2")

	input := ReadFile("day23-input.txt")
	inputList := strings.Split(input, "\n")

	fmt.Println(inputList)
	registers := make(map[string]int64)
	registers["a"] = 1
	fmt.Println(registers)

	handleInstructions(inputList, registers)
	fmt.Println(registers)
}

func DayTwentyThreePartTwoNative() {

	start := (57 * 100) + 100000
	end := start + 17000

	fmt.Println("Start:", start)
	fmt.Println("End:", end)

	notPrime := 0

	for i:= start; i <= end ; i = i+17 {
		fmt.Println("Trying:", i)

		for j:=2; j<i;j++ {

			if i % j == 0 {
				notPrime++
				break
			}
		}
	}

	fmt.Println("Non-prime count:", notPrime)
}
