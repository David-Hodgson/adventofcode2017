package adventofcode2017

import (
	"fmt"
	"strings"
)

const IP_DOWN int = 0
const IP_UP int = 1
const IP_LEFT int = 2
const IP_RIGHT int = 3

func getNextIPDirection(rows []string, x,y, fromDirection int) int {

	//UP
	if (fromDirection != IP_DOWN && y > 0 && x < len(rows[y-1])) {
		upChar := string(rows[y-1][x])
		if upChar != " " {
			return IP_UP
		}
	}

	if (fromDirection != IP_UP && y < len(rows)-1) {
		downChar := string(rows[y+1][x])
		if downChar != " " {
			return IP_DOWN
		}
	}

	if (fromDirection != IP_RIGHT && x > 0) {
		leftChar := string(rows[y][x-1])
		if leftChar != " " {
			return IP_LEFT
		}
	}

	if (fromDirection != IP_LEFT && x < len(rows[y])) {
		rightChar := string(rows[y][x+1])
		if rightChar != " " {
			return IP_RIGHT
		}
	}

	return IP_RIGHT
}

func followIpRoute(rows []string) int {
	var x,y int
	direction := IP_DOWN

	//Find starting pos
	for j:= 0 ; j< len(rows	[0]); j++ {

		if (string(rows[0][j]) == "|") {
			//fmt.Println("Found start at ", j, ",0")
			x = j
		}
	}

	fmt.Println("Starting at",x,",",y,"moving",direction)

	finished := false
	steps := 0
	for ;;{

		steps++
		switch direction {
		case IP_DOWN:
			y++
			break
		case IP_UP:
			y--
			break
		case IP_LEFT:
			if x>0 {
				x--
				break
			} else {
				finished = true
				break
			}
		case IP_RIGHT:
			x++
			break
		}

		current := string(rows[y][x])

		if current == "+" {
			//direction change
			//fmt.Println("Direction Change",x,",",y)

			direction = getNextIPDirection(rows,x,y, direction)

		} else if (current != "-" && current != "|" && current != " ") {
			fmt.Print(current)
		} else if current == " "{
			finished = true
			break
		}

		if finished {
			break
		}
	}

	return steps
}

func DayNineteenExample() {


	fmt.Println("Day 19 - Example")

	input := "     |\n     |  +--+\n     A  |  C\n F---|----E|--+\n     |  |  |  D\n     +B-+  +--+"

	rows := strings.Split(input, "\n")

	for i:= 0; i< len(rows); i++ {
		fmt.Println(rows[i])
	}

	stepCount := followIpRoute(rows)
	fmt.Println()
	fmt.Println("Steps:", stepCount)
}

func DayNineteenPartOne() {


	fmt.Println("Day 19 - Part One")

	input := ReadFile("day19-input.txt")
	rows := strings.Split(input, "\n")


	stepCount := followIpRoute(rows)
	fmt.Println()
	fmt.Println("Steps:", stepCount)
}


