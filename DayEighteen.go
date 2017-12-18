package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)


func getRegisterValue(registerValue string, registers map[string]int64) int64 {
	value,err := strconv.Atoi(registerValue)

	if err == nil {
		return int64(value)
	}

	return registers[registerValue]

}

func processCommands(commandList []string, registers map[string]int64) {

	var lastSound int64
	for i:= 0 ; i<len(commandList) ; {
		fmt.Println(commandList[i])
		commandParts := strings.Split(commandList[i], " ")

		switch commandParts[0] {
		case "set":
			register := commandParts[1]
			value := getRegisterValue(commandParts[2],registers)
			registers[register] = int64(value)
			break
		case "add":
			register := commandParts[1]
			value := getRegisterValue(commandParts[2], registers)
			registers[register] = registers[register] + value
			break
		case "mul":
			register := commandParts[1]
			value := getRegisterValue(commandParts[2], registers)
			registers[register] = registers[register] * value
			break
		case "mod":
			register := commandParts[1]
			value := getRegisterValue(commandParts[2], registers)
			registers[register] = registers[register] % value
			break
		case "snd":
			lastSound = getRegisterValue(commandParts[1], registers)
			break
		case "rcv":
			register := getRegisterValue(commandParts[1], registers)
			if register != 0 {
				i = len(commandList)
			}
			break
		case "jgz":
			value := getRegisterValue(commandParts[1], registers)
			jumpvalue := getRegisterValue(commandParts[2], registers)

			if value > 0 {
				i--
				i = i + int(jumpvalue)
			}
			break
		default:
			fmt.Println(commandParts[0])
		}

		fmt.Println(registers)
		i++
	}

	fmt.Println("Last Sound:",lastSound)
}

type programState struct {
	currentPosition int64
	waiting bool
	registers map[string]int64
	sentValue []int64
	sentCount int
}

func processCommand(commandList []string, state, otherState *programState)  {

	commandParts := strings.Split(commandList[state.currentPosition], " ")

	var commandInc int64 = 1
	switch commandParts[0] {
	case "set":
		register := commandParts[1]
		value := getRegisterValue(commandParts[2],state.registers)
		state.registers[register] = int64(value)
		break
	case "add":
		register := commandParts[1]
		value := getRegisterValue(commandParts[2], state.registers)
		state.registers[register] = state.registers[register] + value
		break
	case "mul":
		register := commandParts[1]
		value := getRegisterValue(commandParts[2], state.registers)
		state.registers[register] = state.registers[register] * value
		break
	case "mod":
		register := commandParts[1]
		value := getRegisterValue(commandParts[2], state.registers)
		state.registers[register] = state.registers[register] % value
		break
	case "snd":
		valueToSend := getRegisterValue(commandParts[1], state.registers)
		state.sentValue = append(state.sentValue, valueToSend)
		state.sentCount++
		break
	case "rcv":
		if len(otherState.sentValue) > 0 {
			state.waiting = false
			//We have another value

			state.registers[commandParts[1]] = otherState.sentValue[0]
			otherState.sentValue = otherState.sentValue[1:]

		} else {
			state.waiting = true
			commandInc = 0
		}
		break
	case "jgz":
		value := getRegisterValue(commandParts[1], state.registers)
		jumpvalue := getRegisterValue(commandParts[2], state.registers)

		if value > 0 {
			commandInc = jumpvalue
		}
		break
	default:
		fmt.Println(commandParts[0])
	}

	state.currentPosition += commandInc

}


func DayEighteenExample() {


	fmt.Println("Day 18 - Example")

	input := "set a 1\nadd a 2\nmul a a\nmod a 5\nsnd a\nset a 0\nrcv a\njgz a -1\nset a 1\njgz a -2"

	inputList := strings.Split(input, "\n")

	registers := make(map[string]int64)

	processCommands(inputList, registers)

	fmt.Println(registers)
}




func DayEighteenPartOne() {


	fmt.Println("Day 18 - Part One")

	input := ReadFile("day18-input.txt")

	inputList := strings.Split(input, "\n")

	registers := make(map[string]int64)

	processCommands(inputList, registers)

	fmt.Println(registers)
}

func DayEighteenPartTwo() {


	fmt.Println("Day 18 - Part Two")

	input := ReadFile("day18-input.txt")

	inputList := strings.Split(input, "\n")

	p0Registers := make(map[string]int64)
	p0Registers["p"] = 0
	p1Registers := make(map[string]int64)
	p1Registers["p"] = 1

	//processCommands(inputList, registers)
	p0State := programState{currentPosition:0,registers:p0Registers}
	p1State := programState{currentPosition:0,registers:p1Registers}

	//p0State.waiting = true
	//p1State.waiting = true;

	for ;; {
		processCommand(inputList,&p0State,&p1State)
		processCommand(inputList,&p1State,&p0State)


		if p0State.waiting && p1State.waiting {
			break
		}
	}

	fmt.Println("P0 state:", p0State)
	fmt.Println("p1 state:", p1State)

}