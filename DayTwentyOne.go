package adventofcode2017

import (
	"fmt"
	"strings"
)

type enchancementRule struct {
	inputMatrix [][]bool
	outputMatrix [][]bool
}

func parseEnchancementRule(input string) enchancementRule{
	ruleParts := strings.Split(input, "=>")

	sourceMatrixString := strings.Trim(ruleParts[0], " ")
	destinationMatrixString := strings.Trim(ruleParts[1], " ")

	sourceMatrixString = strings.Replace(sourceMatrixString, "/", "\n",-1)
	destinationMatrixString = strings.Replace(destinationMatrixString, "/", "\n",-1)

	sourceMatrix := makeMatrix(sourceMatrixString)
	destinationMatrix := makeMatrix(destinationMatrixString)

	return enchancementRule{sourceMatrix, destinationMatrix}
}

func printMatrix(matrix [][]bool) {

	for i := 0; i<len(matrix);i++ {

		for j:=0 ; j< len(matrix[i]); j++ {
			if matrix[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func rotateMatrix(matrix [][]bool) [][]bool {
	matrixSize := len(matrix)
	outputMatrix := make([][]bool, matrixSize)

	for i := 0; i < matrixSize ; i++  {
		outputMatrix[i] = make([]bool, matrixSize)

		for j := 0; j < matrixSize; j++  {
			outputMatrix[i][j] = matrix[matrixSize-j-1][i]
		}
	}

	return outputMatrix
}

func flipMatrixVertical(matrix [][]bool) [][]bool {
	matrixSize := len(matrix)
	outputMatrix := make([][]bool, matrixSize)

	for i := 0; i < matrixSize ; i++  {
		outputMatrix[i] = make([]bool, matrixSize)

		for j := 0; j < matrixSize; j++  {
			outputMatrix[i][j] = matrix[matrixSize-i -1][j]
		}
	}

	return outputMatrix
}

func flipMatrixHorizontal(matrix [][]bool) [][]bool {

	matrixSize := len(matrix)
	outputMatrix := make([][]bool, matrixSize)

	for i := 0; i < matrixSize ; i++  {
		outputMatrix[i] = make([]bool, matrixSize)

		for j := 0; j < matrixSize; j++  {
			outputMatrix[i][j] = matrix[i][matrixSize -j - 1]
		}
	}

	return outputMatrix
}

func makeMatrix(stringDescription string) [][]bool {

	startingMatrixDescriptions := strings.Split(stringDescription, "\n")

	matrix := make([][]bool, len(startingMatrixDescriptions))

	for i:= 0 ; i<len(matrix);i++ {

		matrix[i] = make([]bool, len(startingMatrixDescriptions[i]));

		for j:=0; j< len(startingMatrixDescriptions[i]); j++ {

			if startingMatrixDescriptions[i][j:j+1] == "#"{
				matrix[i][j] = true
			}
		}
	}

	return matrix

}

func areMatrixEqual(m1,m2[][]bool) bool {

	if len(m1) != len(m2) {
		return false
	}

	matching := true

	for i:=0 ; i<len(m1) && matching;i++ {
		for j := 0;j<len(m1) && matching; j++ {
			if m1[i][j] != m2[i][j] {
				matching = false
			}
		}
	}

	return matching
}

func buildMatrix(size int) [][]bool {
	outputMatrix := make([][]bool,size)

	for i:=0 ; i<size ; i++ {
		outputMatrix[i] = make([]bool,size)
	}

	return outputMatrix
}

func applyRule(inputMatrix [][]bool, rules []enchancementRule) [][]bool {

	matrixCombos := make([][][]bool, 8)

	matrixCombos[0] = inputMatrix
	matrixCombos[1] = rotateMatrix(matrixCombos[0])
	matrixCombos[2] = rotateMatrix(matrixCombos[1])
	matrixCombos[3] = rotateMatrix(matrixCombos[2])
	matrixCombos[4] = flipMatrixHorizontal(matrixCombos[0])
	matrixCombos[5] = rotateMatrix(matrixCombos[4])
	matrixCombos[6] = rotateMatrix(matrixCombos[5])
	matrixCombos[7] = rotateMatrix(matrixCombos[6])

	for i := 0; i < len(rules); i++ {
		for j := 0 ; j < len(matrixCombos) ;j++ {
			if areMatrixEqual(rules[i].inputMatrix, matrixCombos[j]) {
				return rules[i].outputMatrix
			}
		}

	}
	return inputMatrix
}

func enchanceMatrix(matrix [][]bool, rules []enchancementRule) [][]bool {

	chunkSize := -1

	if len(matrix) % 2 == 0 {
		chunkSize = 2
	} else 	if len(matrix) % 3 == 0 {
		chunkSize = 3
	}

	if chunkSize < 0 {
		return matrix
	}

	chunks := len(matrix) / chunkSize
	outputMatrix := buildMatrix((chunkSize+1) * chunks)

	for i:=0 ;i<chunks; i++ {
		for j:=0 ;j<chunks; j++ {

			//create sub matrix
			subMatrix := make([][]bool, chunkSize)

			for x:=0; x<chunkSize;x++ {
				subMatrix[x] = make([]bool, chunkSize)
				for y:=0;y<chunkSize;y++ {
					subMatrix[x][y] = matrix[x+(i*chunkSize)][y+(j*chunkSize)]
				}
			}

			subMatrix = applyRule(subMatrix, rules)

			for m :=0 ; m < len(subMatrix); m++ {
				for n:=0; n < len(subMatrix[m]); n++ {
					outputMatrix[(i*(chunkSize+1))+m][(j*(chunkSize+1))+n] = subMatrix[m][n]
				}
			}

		}
	}

	return outputMatrix
}


func DayTwentyOneExample() {

	fmt.Println("Day 20 - Example")

	matrix := makeMatrix(".#.\n..#\n###")

	printMatrix(matrix)

	rules := make([]enchancementRule,2)
	rules[0] = parseEnchancementRule("../.# => ##./#../...")
	rules[1] = parseEnchancementRule(".#./..#/### => #..#/..../..../#..#")

	for i:=0 ;i<3;i ++ {
		matrix = enchanceMatrix(matrix,rules)
		printMatrix(matrix)
	}

	pixelCount := 0

	for i :=0; i < len(matrix);i++ {
		for j:=0 ; j< len(matrix[i]); j++ {
			if matrix[i][j] {
				pixelCount++
			}
		}
	}

	fmt.Println("Lit Pixel Count:", pixelCount)
}

func DayTwentyOnePartOne() {

	fmt.Println("Day 20 - Part One")

	matrix := makeMatrix(".#.\n..#\n###")

	printMatrix(matrix)

	rulesInput := strings.Split(ReadFile("day21-input.txt"), "\n")

	rules := make([]enchancementRule,len(rulesInput))
	for x:= 0; x <len(rulesInput); x++ {
		rules[x] = parseEnchancementRule(rulesInput[x])
	}

	for i:=0 ;i<5 ;i ++ {
		matrix = enchanceMatrix(matrix,rules)
		printMatrix(matrix)
	}

	pixelCount := 0

	for i :=0; i < len(matrix);i++ {
		for j:=0 ; j< len(matrix[i]); j++ {
			if matrix[i][j] {
				pixelCount++
			}
		}
	}

	fmt.Println("Lit Pixel Count:", pixelCount)
}

func DayTwentyOnePartTwo() {

	fmt.Println("Day 20 - Part Two")

	matrix := makeMatrix(".#.\n..#\n###")

	printMatrix(matrix)

	rulesInput := strings.Split(ReadFile("day21-input.txt"), "\n")

	rules := make([]enchancementRule,len(rulesInput))
	for x:= 0; x <len(rulesInput); x++ {
		rules[x] = parseEnchancementRule(rulesInput[x])
	}

	for i:=0 ;i<18 ;i ++ {
		matrix = enchanceMatrix(matrix,rules)
	}

	pixelCount := 0

	for i :=0; i < len(matrix);i++ {
		for j:=0 ; j< len(matrix[i]); j++ {
			if matrix[i][j] {
				pixelCount++
			}
		}
	}

	fmt.Println("Lit Pixel Count:", pixelCount)
}

