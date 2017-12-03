package adventofcode2017

import (
	"fmt"
)

const RIGHT=0
const UP = 1
const LEFT = 2
const DOWN = 3

func getNextDirection(currentDirection int) int {

	switch currentDirection {
	case RIGHT:
		return UP
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	case DOWN:
		return RIGHT
	default:
		return UP
	}
}

func move (currentDirection,x,y int) (newX,newY int) {
	switch currentDirection {
	case RIGHT:
		return x+1,y
	case UP:
		return x,y+1
	case LEFT:
		return x-1,y
	case DOWN:
		return x,y-1
	default:
		return x,y
	}

}

func populateGrid(maxNumber int, height,width int) [][]int {
	step := 1

	currentDirection := RIGHT
	currentStep := step

	stepsTillStepChange := 2

	var grid [][]int
	grid = make([][]int,height)

	for i := 0; i<len(grid);i++ {

		grid[i] = make([]int, width)

	}

	x := width /2
	y := height /2

	for i:=1;i<=maxNumber;i++ {
		//fmt.Println("i: ", i, " Current Direction: ", currentDirection, "Step: ", step, "CurrentStep: ", currentStep)
		//fmt.Println("Setting",i, "into",x,",",y)
		grid[y][x] = i

		if i==maxNumber {
			fmt.Println("",x,",",y)
		}
		x,y = move(currentDirection, x,y)
		currentStep--

		if (currentStep == 0) {
			//fmt.Println("Chaning direction")
			stepsTillStepChange--
			if stepsTillStepChange == 0 {
				step++
				stepsTillStepChange = 2
			}
			currentStep = step
			currentDirection = getNextDirection(currentDirection)
		}


	}

	//fmt.Println("",x,":",y)
	return grid
}

func populateGridUsingAdjacentCells(maxNumber int, height,width int) [][]int {
	step := 1

	currentDirection := RIGHT
	currentStep := step

	stepsTillStepChange := 2

	var grid [][]int
	grid = make([][]int,height)

	for i := 0; i<len(grid);i++ {

		grid[i] = make([]int, width)

	}

	x := width /2
	y := height /2

	for i:=1;i<=maxNumber;i++ {
		//fmt.Println("i: ", i, " Current Direction: ", currentDirection, "Step: ", step, "CurrentStep: ", currentStep)
		//fmt.Println("Setting",i, "into",x,",",y)
		//grid[y][x] = i
		if i== 1 {
			grid[y][x] = i
		} else {
			grid[y][x] = grid[y - 1][x - 1] + grid[y - 1][x] + grid[y - 1][x + 1] +
				grid[y][x - 1] + grid[y][x + 1] +
				grid[y + 1][x - 1] + grid[y + 1][x] + grid[y + 1][x + 1]
		}

		if grid[y][x] > maxNumber {
			fmt.Println("Max Number is", grid[y][x])
			break
		}

		if i==maxNumber {
			fmt.Println("",x,",",y)
		}
		x,y = move(currentDirection, x,y)
		currentStep--

		if (currentStep == 0) {
			//fmt.Println("Chaning direction")
			stepsTillStepChange--
			if stepsTillStepChange == 0 {
				step++
				stepsTillStepChange = 2
			}
			currentStep = step
			currentDirection = getNextDirection(currentDirection)
		}


	}

	//fmt.Println("",x,":",y)
	return grid
}


func getGridHeightAndWidth(maxNumber int) (height,width int) {
	step := 1

	currentDirection := RIGHT
	currentStep := step

	stepsTillStepChange := 2

	width = step
	height = step

	for i:=1;i<=maxNumber;i++ {
		//fmt.Println("i: ", i, " Current Direction: ", currentDirection, "Step: ", step, "CurrentStep: ", currentStep, "StepsTillChange", stepsTillStepChange)
		currentStep--

		if (currentStep == 0) {
			//fmt.Println("Chaning direction")
			stepsTillStepChange--
			if stepsTillStepChange == 0 {
				step++
				stepsTillStepChange = 2
			}
			currentStep = step
			currentDirection = getNextDirection(currentDirection)
			width = step
			height = step
		}

	}

	return height,width
}

func DayThreeExample() {

	fmt.Println("Day Three Example")

	max := 23

	height,width := getGridHeightAndWidth(max)

	fmt.Println("Width: ", width)
	fmt.Println("Height: ", height)

	fmt.Println("Start x:", width / 2)
	fmt.Println("Start y:", height / 2)


	grid := populateGrid(max, height,width)

	fmt.Println(grid)
}

func DayThreePartOne() {

	fmt.Println("Day Three Part One")

	max := 265149

	height,width := getGridHeightAndWidth(max)

	fmt.Println("Width: ", width)
	fmt.Println("Height: ", height)

	fmt.Println("Start x:", width / 2)
	fmt.Println("Start y:", height / 2)


	grid := populateGrid(max, height,width)

	fmt.Println(len(grid))
	//fmt.Println(grid)
}

func DayThreePartTwo() {

	fmt.Println("Day Three Part One")

	max := 265149

	height,width := getGridHeightAndWidth(max)

	fmt.Println("Width: ", width)
	fmt.Println("Height: ", height)

	fmt.Println("Start x:", width / 2)
	fmt.Println("Start y:", height / 2)


	grid := populateGridUsingAdjacentCells(max, height,width)

	fmt.Println(len(grid))
	//fmt.Println(grid)
}