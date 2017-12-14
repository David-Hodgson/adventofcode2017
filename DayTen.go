package adventofcode2017

import (
	"fmt"
	//"strconv"
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

func convertSparseHashToDense(loop []int) []int {


	dense := make([]int, len(loop)/16)

	offset := 0;
	for i := 0 ; i < len(dense); i++ {

		dense[i] = loop[offset] ^ loop[offset+1] ^ loop[offset+2] ^ loop[offset+3] ^ loop[offset+4] ^
			loop[offset+5]^ loop[offset+6] ^ loop[offset+7] ^ loop[offset+8] ^ loop[offset+9] ^
			^ loop[offset+10] ^ loop[offset+11] ^ loop[offset+12] ^ loop[offset+13] ^
			^ loop[offset+14] ^ loop[offset+15]

		offset += 16
	}

	return dense
}

func toHexString(loop []int) string {

	var hex string
	for i := 0 ; i< len(loop); i++ {
		h := fmt.Sprintf("%02x", loop[i])
		hex += h
	}

	return hex

}

var(
position int = 0
skip int = 0
)

func process(loop []int, lengths []int) {

	for i := 0; i <len(lengths) ; i++ {
		twist(loop, position, lengths[i])
		position += lengths[i] + skip

		if position >= len(loop) {
			position = position % len(loop)
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

	position =0
	skip = 0
	process(loop, inputs)
	fmt.Println("Item[0]:", loop[0])
	fmt.Println("Item[1]:", loop[1])

	fmt.Println("Result:", (loop[0]*loop[1]))
}

func KnotHash(input string) string {

	loop := buildLoop(256)

	inputs := make([]int, len(input));

	for i := 0 ; i < len(input) ; i++ {
		inputs[i] = int(input[i])
	}
	extras := [] int {17, 31, 73, 47, 23}

	mergedlengths := make([]int, (len(inputs) + len(extras)))

	for i :=0; i<len(inputs);i++ {

		mergedlengths[i] = inputs[i]
	}

	extraPos :=0
	for i := len(inputs); i < len(mergedlengths); i++ {
		mergedlengths[i] = extras[extraPos]
		extraPos++
	}

	position = 0
	skip = 0

	for i :=0 ;i < 64 ; i++ {
		process(loop, mergedlengths)
	}

	dense := convertSparseHashToDense(loop)

	hex := toHexString(dense)

	return hex
}

func DayTenPartTwo() {

	fmt.Println("Day Ten - Part Two")

	inputString  := "147,37,249,1,31,2,226,0,161,71,254,243,183,255,30,70"

	hex := KnotHash(inputString)

	fmt.Println("Knot Hex:", hex)

}

