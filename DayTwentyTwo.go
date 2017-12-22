package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

const NORTH = 1
const SOUTH = 2
const EAST = 3
const WEST = 4

const CLEAN = 0
const INFECTED = 1
const CLEANED = 2

func turnRight(currentDir int) int {

	switch(currentDir) {
		case NORTH:
			return EAST
		case EAST:
			return SOUTH
		case SOUTH:
			return WEST
		case WEST:
			return NORTH
		default:
			return NORTH
	}
}

func turnLeft(currentDir int) int {

	switch(currentDir) {
		case NORTH:
			return WEST
		case WEST:
			return SOUTH
		case SOUTH:
			return EAST
		case EAST:
			return NORTH
		default:
			return NORTH
	}
}

func moveVirusCarrier(x,y,currentDirection int) (newX,newY int) {

	switch(currentDirection) {
		case NORTH:
			newX = x
			newY = y+1
			break
		case EAST:
			newX = x+1
			newY = y
			break
		case SOUTH:
			newX = x
			newY = y-1
			break
		case WEST:
			newX = x-1
			newY = y
			break
		default:
			newX = x
			newY = y
	}

	return newX,newY
}

var infectionCount int = 0

func virusCarrierDoWork(x,y,currentDirection int,grid map[string]int) (int,int, int) {

	//if current node is infect - turn right else turn left
	currentPosString := xyToString(x,y)
	currentState := grid[currentPosString]

	if currentState == INFECTED {
		currentDirection = turnRight(currentDirection)
	} else {
		currentDirection = turnLeft(currentDirection)
	}

	//if current node is clean - node becomes infected else becomes cleaned
	if currentState != INFECTED {
		grid[currentPosString] = INFECTED
		infectionCount++
		fmt.Println("Infecting. Count:",infectionCount)
	} else {
		grid[currentPosString] = CLEANED
	}
	//move one node forward
	x,y = moveVirusCarrier(x,y,currentDirection)
	return x,y,currentDirection
}

func xyToString(x,y int) string {

	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func DayTwentyTwoExample() {

	fmt.Println("Day 22 - Example")

	x,y := 0,0
	dir := NORTH

	fmt.Println("x:",x,"y:",y)
	fmt.Println("Current Direction:",dir)

	grid := make(map[string]int)

	inputGrid := "..#\n#..\n..."

	gridRows := strings.Split(inputGrid, "\n")

	gridSize := len(gridRows)
	fmt.Println("Grid Size:", gridSize)

	centerX := gridSize /2
	centerY := centerX

	fmt.Println("center point, x:", centerX, "y:", centerY)

	for gridY := 0 ; gridY < len(gridRows) ; gridY++ {
		fmt.Println("Grid Row:", gridY)
		fmt.Println("Gird Row - center:", (gridY -centerY)*-1)
		convertedY := (gridY-centerY)*-1
		for gridX := 0 ; gridX < len(gridRows[gridY]); gridX++ {
			convertedX := (gridX - centerX)
			if string(gridRows[gridY][gridX]) == "#" {
				grid[xyToString(convertedX,convertedY)] = INFECTED
			}
		}
	}

	fmt.Println(grid)

	for i:= 0 ; i< 10000; i++ {
		x,y,dir = virusCarrierDoWork(x,y,dir,grid)
	}
	fmt.Println("x:",x,"y:",y)
	fmt.Println("Current Direction:",dir)
	fmt.Println("Infection Count:", infectionCount)
}


