package adventofcode2017

import (
	"fmt"
)

const GROUP =1
const GARBAGE = 2

func processStream(streamData string) int {

	var state int

	groupDepth := 0
	total := 0

	for i := 0 ; i < len(streamData) ; i++ {

		current := string(streamData[i])



		if state == GARBAGE {
			if current == ">" {
				state = GROUP
			}

			if current == "!" {
				i++
			}
		} else {

			if current == "{" {
				state = GROUP
				groupDepth++
				total += groupDepth
			}

			if current == "}" {
				state = GROUP
				groupDepth--
			}

			if current == "<" {
				state = GARBAGE

			}
		}

	}

	return total
}

func DayNineExample() {

	fmt.Println("Day Nine - Example")

	input := "{}"
	fmt.Println("Total:", processStream(input))

	input = "{{{}}}"
	fmt.Println("Total:", processStream(input))

	input = "{{},{}}"
	fmt.Println("Total:", processStream(input))

	input = "{{{},{},{{}}}}"
	fmt.Println("Total:", processStream(input))

	input = "{<a>,<a>,<a>,<a>}"
	fmt.Println("Total:", processStream(input))

	input = "{{<ab>},{<ab>},{<ab>},{<ab>}}"
	fmt.Println("Total:", processStream(input))

	input = "{{<!!>},{<!!>},{<!!>},{<!!>}}"
	fmt.Println("Total:", processStream(input))

	input = "{{<a!>},{<a!>},{<a!>},{<ab>}}"
	fmt.Println("Total:", processStream(input))
}

func DayNinePartOne() {

	input := ReadFile("day9-input.txt")
	fmt.Println("Day Nine - Part One - Total:", processStream(input))
}




