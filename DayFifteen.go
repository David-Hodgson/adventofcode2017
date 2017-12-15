package adventofcode2017

import (
	"fmt"
)

const MAGIC_DIVISOR = 2147483647

func runGenerator(startValue, factor int) int {
	return (startValue * factor) % MAGIC_DIVISOR

}

func runGeneratorPickey(startValue, factor int, picky int) int {
	value := (startValue * factor) % MAGIC_DIVISOR

	if value % picky == 0 {
		return  value
	} else {
		return runGeneratorPickey(value, factor, picky)
	}
}

func DayFifteenExample() {

	fmt.Println("Day 15 - Example")

	genAValue := 65
	genAFactor := 16807

	genBValue := 8921
	genBFactor := 48271

	score := 0
	for i := 0 ; i < 5 ; i++ {
		genAValue = runGenerator(genAValue, genAFactor)
		genBValue = runGenerator(genBValue, genBFactor)

		binaryAValue := fmt.Sprintf("%016b", genAValue)
		binaryBValue := fmt.Sprintf("%016b", genBValue)

		binaryAValue = binaryAValue[len(binaryAValue)-16:]
		binaryBValue = binaryBValue[len(binaryBValue)-16:]

		if binaryAValue == binaryBValue {
			fmt.Println("Match at",i	)
			score++
		}

	}

	fmt.Println("Score:", score)

}


func DayFifteenPartOne() {

	fmt.Println("Day 15 - Part One")

	genAValue := 512
	genAFactor := 16807

	genBValue := 191
	genBFactor := 48271

	score := 0
	for i := 0; i < 40000000; i++ {
		genAValue = runGenerator(genAValue, genAFactor)
		genBValue = runGenerator(genBValue, genBFactor)

		binaryAValue := fmt.Sprintf("%016b", genAValue)
		binaryBValue := fmt.Sprintf("%016b", genBValue)

		binaryAValue = binaryAValue[len(binaryAValue)-16:]
		binaryBValue = binaryBValue[len(binaryBValue)-16:]

		if binaryAValue == binaryBValue {
			score++
		}

	}

	fmt.Println("Score:", score)

}

func DayFifteenExample2() {

	fmt.Println("Day 15 - Example 2")

	genAValue := 65
	genAFactor := 16807

	genBValue := 8921
	genBFactor := 48271

	score := 0
	for i := 0 ; i < 5000000 ; i++ {
		genAValue = runGeneratorPickey(genAValue, genAFactor, 4)
		genBValue = runGeneratorPickey(genBValue, genBFactor,8)

		binaryAValue := fmt.Sprintf("%016b", genAValue)
		binaryBValue := fmt.Sprintf("%016b", genBValue)

		binaryAValue = binaryAValue[len(binaryAValue)-16:]
		binaryBValue = binaryBValue[len(binaryBValue)-16:]

		if binaryAValue == binaryBValue {
			score++
		}

	}

	fmt.Println("Score:", score)

}

func DayFifteenPartTwo() {

	fmt.Println("Day 15 - Part Two")

	genAValue := 512
	genAFactor := 16807

	genBValue := 191
	genBFactor := 48271

	score := 0
	for i := 0; i < 5000000; i++ {
		genAValue = runGeneratorPickey(genAValue, genAFactor, 4)
		genBValue = runGeneratorPickey(genBValue, genBFactor,8)

		binaryAValue := fmt.Sprintf("%016b", genAValue)
		binaryBValue := fmt.Sprintf("%016b", genBValue)

		binaryAValue = binaryAValue[len(binaryAValue)-16:]
		binaryBValue = binaryBValue[len(binaryBValue)-16:]

		if binaryAValue == binaryBValue {
			score++
		}

	}

	fmt.Println("Score:", score)

}