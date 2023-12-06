package day1

import (
	"aoc2023/common"
	"fmt"
	"strconv"
)

func isWordCorrect(str string, num int) bool {
	var numbers [10]string = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	return str == numbers[num]
}

func whichPossibleStartingLetter(b byte) []int {
	str := string(b)
	var startingLetters [10]string = [10]string{"z", "o", "t", "t", "f", "f", "s", "s", "e", "n"}
	result := []int{}
	for i := 0; i <= 9; i++ {
		if str == startingLetters[i] {
			result = append(result, i)
		}
	}
	return result
}

func whichPossibleEndingLetter(b byte) []int {
	str := string(b)
	var endingLetters [10]string = [10]string{"o", "e", "o", "e", "r", "e", "x", "n", "t", "e"}
	result := []int{}
	for i := 0; i <= 9; i++ {
		if str == endingLetters[i] {
			result = append(result, i)
		}
	}
	return result
}

func amountOfLettersInNumber(num int) int {
	switch num {
	case 0:
		return 4
	case 1:
		return 3
	case 2:
		return 3
	case 3:
		return 5
	case 4:
		return 4
	case 5:
		return 4
	case 6:
		return 3
	case 7:
		return 5
	case 8:
		return 5
	case 9:
		return 4
	default:
		panic("invalid input on amountOfLettersInNumber()")
	}
}

func searchLeftmostNumeric(str string) string {
	result := ""

	for i := 0; i < len(str) && result == ""; i++ {
		if common.IsNumber(str[i]) {
			result = string(str[i])
		} else {
			possibleNumbers := whichPossibleStartingLetter(str[i])
			if len(possibleNumbers) > 0 {
				for j := 0; j < len(possibleNumbers) && result == ""; j++ {
					possibleNumber := possibleNumbers[j]
					amountToSlice := amountOfLettersInNumber(possibleNumber)
					if i+amountToSlice < len(str) {
						possibleWord := str[i : amountToSlice+i]
						if isWordCorrect(possibleWord, possibleNumber) {
							result = strconv.Itoa(possibleNumber)
						}
					}
				}
			}
		}
	}

	return result
}

func searchRightmostNumeric(str string) string {
	result := ""

	for i := len(str) - 1; i >= 0 && result == ""; i-- {
		if common.IsNumber(str[i]) {
			result = string(str[i])
		} else {
			possibleNumbers := whichPossibleEndingLetter(str[i])
			if len(possibleNumbers) > 0 {
				for j := 0; j < len(possibleNumbers) && result == ""; j++ {
					possibleNumber := possibleNumbers[j]
					amountToSlice := amountOfLettersInNumber(possibleNumber)
					if i-amountToSlice >= 0 {
						possibleWord := str[i-amountToSlice+1 : i+1]
						if isWordCorrect(possibleWord, possibleNumber) {
							result = strconv.Itoa(possibleNumber)
						}
					}
				}
			}
		}
	}

	return result
}

func Part2() {
	lines := common.FileToSlice("./day1/input.txt")
	total := 0
	for idx, currStr := range lines {
		foundCode := searchLeftmostNumeric(currStr) + searchRightmostNumeric(currStr)
		value, err := strconv.Atoi(foundCode)
		if err != nil {
			fmt.Printf("Error at line %d", idx)
			panic(err)
		}
		total += value
	}
	fmt.Printf("The result is %d\n", total)
}
