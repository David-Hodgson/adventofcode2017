package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func ConvertStringToSpreadSheet(input string) [][]int {

	var spreadSheet  [][]int

	rows := strings.Split(input, "\n")

	spreadSheet = make([][]int, len(rows))

	for i := 0 ;i < len(rows) ;i++ {
		fmt.Println(rows[i])

		cells := strings.Split(rows[i], "\t")
		spreadSheet[i] = make([]int, len(cells))
		for j:= 0; j< len(cells); j++ {
			spreadSheet[i][j],_ = strconv.Atoi(cells[j])
		}
	}

	return  spreadSheet
}

func CalculateCheckSum(spreadsheet [][]int) int {

	checksum := 0

	for i:= 0; i< len(spreadsheet); i++ {

		sort.Ints(spreadsheet[i]);
		top := spreadsheet[i][len(spreadsheet[i])-1]
		bottom := spreadsheet[i][0]

		if (top > bottom) {
			checksum += top - bottom
		} else {
			checksum += bottom - top
		}
	}

	return checksum
}

func CalculateEvenCheckSum(spreadsheet [][]int) int {

	checksum := 0

	for rowNumber:= 0; rowNumber< len(spreadsheet); rowNumber++ {

		row := spreadsheet[rowNumber]

		for i := 0 ;i < len(row); i++ {

			for j:=0 ;j < len(row) ; j++ {

				if i != j {

					if row[i] % row[j] == 0 {
						fmt.Println("",row[i], "%", row[j], "is 0")
						checksum += row[i] / row[j]
						fmt.Println(checksum)
					}
				}

			}
		}
	}

	return checksum
}

func DayTwoExample () {

	input := "5 1 9 5\n7 5 3\n2 4 6 8";

	sheet := ConvertStringToSpreadSheet(input)

	fmt.Println(sheet)

	fmt.Println(CalculateCheckSum(sheet))
}

func DayTwoPartOne() {

	input := ReadFile("day2-input.txt")

	sheet := ConvertStringToSpreadSheet(input)

	//fmt.Println(sheet)

	fmt.Println(CalculateCheckSum(sheet))
}

func DayTwoPartTwo() {

	input := ReadFile("day2-input.txt")

	sheet := ConvertStringToSpreadSheet(input)

	//fmt.Println(sheet)

	fmt.Println(CalculateEvenCheckSum(sheet))
}