package adventofcode2017

import (
	"fmt"
	"strings"
	"sort"
)


func isValid(input string) bool {

	wordMap := make(map[string]bool)

	words := strings.Split(input, " ");

	for i := 0; i < len(words);i++ {

		word := words[i]
		if _, exists := wordMap[word]; exists {
			return false
		} else {
			wordMap[word] = true
		}

	}
	return true
}


func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isValidAnagram(input string) bool {

	wordMap := make(map[string]bool)

	words := strings.Split(input, " ");

	for i := 0; i < len(words);i++ {

		word := SortString(words[i])

		if _, exists := wordMap[word]; exists {
			return false
		} else {
			wordMap[word] = true
		}

	}
	return true
}

func DayFourExample() {

	fmt.Println("Day Four - Example")

	fmt.Println(isValid("aa bb cc dd ee"))
	fmt.Println(isValid("aa bb cc dd aa"))
	fmt.Println(isValid("aa bb cc dd aaa"))

}

func DayFourPartOne() {
	fmt.Println("Day Four - Example")

	input := ReadFile("day4-input.txt")
	inputStrings := strings.Split(input, "\n")

	validCount := 0

	for i := 0 ; i<len(inputStrings); i++ {

		if len(strings.Trim(inputStrings[i], "")) == 0 {
			continue
		}

		if (isValid(inputStrings[i])) {
			validCount++
		}
	}

	fmt.Println("Valid Count:", validCount)
}

func DayFourPartTwo() {
	fmt.Println("Day Four - Example")

	input := ReadFile("day4-input.txt")
	inputStrings := strings.Split(input, "\n")

	validCount := 0

	for i := 0 ; i<len(inputStrings); i++ {

		if len(strings.Trim(inputStrings[i], "")) == 0 {
			continue
		}

		if (isValidAnagram(inputStrings[i])) {
			validCount++
		}
	}

	fmt.Println("Valid Count:", validCount)
}