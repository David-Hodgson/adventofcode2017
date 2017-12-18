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

	var lastSound int64 = 0
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