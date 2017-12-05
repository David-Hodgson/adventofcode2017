package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)

func processInputToJumpList(input string) []int {

	if strings.HasSuffix(input, "\n") {
		input = input[0:len(input)-1]
	}

	offsets := strings.Split(input, "\n")

	jumps := make([]int,len(offsets))

	for i :=0 ;i<len(offsets);i++ {
		jumps[i],_ = strconv.Atoi(offsets[i])
	}

	return jumps
}

func processJumps(jumpList []int) int {

	step := 0
	currentPos := 0

	for ;currentPos >= 0 && currentPos < len(jumpList) ; {

		value := jumpList[currentPos]

		step++
		//fmt.Println("CurrentPos:", currentPos, "value:", value, "steps:", step)

		jumpList[currentPos] = value + 1
		currentPos += value
	}

	return step
}

func processJumpsV2(jumpList []int) int {

	step := 0
	currentPos := 0

	for ;currentPos >= 0 && currentPos < len(jumpList) ; {

		value := jumpList[currentPos]

		step++
		//fmt.Println("CurrentPos:", currentPos, "value:", value, "steps:", step)

		if value < 3 {
			jumpList[currentPos] = value + 1
		} else {
			jumpList[currentPos] = value - 1
		}
		currentPos += value
	}

	return step
}

func DayFiveExample() {

	input := "0\n3\n0\n1\n-3\n"

	jumps := processInputToJumpList(input)
	fmt.Println(processJumps(jumps))
}

func DayFiveExamplePartTwo() {

	input := "0\n3\n0\n1\n-3\n"

	jumps := processInputToJumpList(input)
	fmt.Println(processJumpsV2(jumps))
}

func DayFivePartOne() {

	input := ReadFile("day5-input.txt")
	jumps := processInputToJumpList(input)
	//fmt.Println(jumps)
	fmt.Println(processJumps(jumps))

}

func DayFivePartTwo() {

	input := ReadFile("day5-input.txt")
	jumps := processInputToJumpList(input)
	//fmt.Println(jumps)
	fmt.Println(processJumpsV2(jumps))

}