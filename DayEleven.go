package adventofcode2017

import (
	"fmt"
	"strings"
)

type point struct {
	x,y int
}

func hexmove(start point, direction string) point {

	x := start.x
	y := start.y
	z := -(x+y)

	switch direction {
	case "ne":
		x++
		y--
		break
	case "n":
		x++
		z--
		break
	case "nw":
		y++
		z--
		break
	case "sw":
		x--
		y++
		break
	case "s":
		x--
		z++
		break
	case "se":
		y--
		z++
		break
	}

	return point{x,y}

}

func getHexDistanceFromOrigin(point point) int {

	x := point.x
	y := point.y
	z := -(x+y)

	return (abs(x) + abs(y) + abs(z)) / 2


}

func abs(value int) int {
	if value < 0 {
		return value * -1
	}

	return value
}

func DayElevenExample() {

	fmt.Println("Day Eleven - Example")

	origin := point{x:0,y:0}

	end := hexmove(hexmove(hexmove(origin, "ne"), "ne"), "ne")

	fmt.Println(end)
	fmt.Println("Distance:", getHexDistanceFromOrigin(end))

	input := "se,sw,se,sw,sw"
	directions := strings.Split(input, ",")

	end = point{0,0}

	for i := 0 ; i < len(directions); i++ {
		end = hexmove(end, directions[i])
	}

	fmt.Println(end)
	fmt.Println("Distance:", getHexDistanceFromOrigin(end))


}

func DayElevenPartOne() {

	fmt.Println("Day Eleven - Part One")

	input := ReadFile("day11-input.txt")
	directions := strings.Split(input, ",")

	end := point{0,0}

	for i := 0 ; i < len(directions); i++ {
		end = hexmove(end, directions[i])
	}

	fmt.Println(end)
	fmt.Println("Distance:", getHexDistanceFromOrigin(end))

}

func DayElevenPartTwo() {

	fmt.Println("Day Eleven - Part Two")

	input := ReadFile("day11-input.txt")
	directions := strings.Split(input, ",")

	end := point{0,0}
	maxDistance := 0

	for i := 0 ; i < len(directions); i++ {
		end = hexmove(end, directions[i])
		distance := getHexDistanceFromOrigin(end)

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Println(end)
	fmt.Println("Distance:", getHexDistanceFromOrigin(end))
	fmt.Println("Max Distance:", maxDistance)

}
