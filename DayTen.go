package adventofcode2017

import (
	"fmt"
)

func buildLoop(length int) []int {

	loop := make([]int, length)

	for i :=0 ; i< length ; i++ {
		loop[i] = i
	}

	return loop
}

func twist(loop []int, position int, length int) {


	//fmt.Println("Twist. position:", position, "length:", length)
	var subList []int

	subList = make([]int, length)

	readPos := position
	for  i := 0; i <length ;i++ {
		if readPos >= len(loop) {
			readPos = 0
		}

		subList[i] = loop[readPos]
		readPos++
	}

	//fmt.Println(subList)

	writePos := position

	for j := len(subList) ; j > 0 ;j-- {
		if writePos >= len(loop) {
			writePos = 0
		}

		loop[writePos] = subList[j-1]
		writePos++
	}

	//fmt.Println(loop)
}

func process(loop []int, lengths []int) {

	position := 0
	skip := 0

	for i := 0; i <len(lengths) ; i++ {
		//fmt.Println(position)
		//fmt.Println(skip)
		twist(loop, position, lengths[i])
		position += lengths[i] + skip

		if position >= len(loop) {
			position = position - len(loop)
		}

		skip++
	}

}

func DayTenExample() {

	fmt.Println("Day Ten - Example")

	loop := buildLoop(5)
	inputs  := []int  {3,4,1,5}

	process(loop, inputs)
	fmt.Println("Item[0]:", loop[0])
	fmt.Println("Item[1]:", loop[1])

	fmt.Println("Result:", (loop[0]*loop[1]))
}

func DayTenPartOne() {

	fmt.Println("Day Ten - Part One")

	loop := buildLoop(256)
	inputs  := []int  {147,37,249,1,31,2,226,0,161,71,254,243,183,255,30,70}

	process(loop, inputs)
	fmt.Println("Item[0]:", loop[0])
	fmt.Println("Item[1]:", loop[1])

	fmt.Println("Result:", (loop[0]*loop[1]))
}


