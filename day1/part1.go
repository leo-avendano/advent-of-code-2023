package day1

import (
	"aoc2023/common"
	"fmt"
	"strconv"
)

func searchNearestNumeric(str string, reverse bool) string {
	result := ""
	if reverse {
		for i := len(str) - 1; i >= 0 && result == ""; i-- {
			if common.IsNumber(str[i]) {
				result = string(str[i])
			}
		}
	} else {
		for i := 0; i < len(str) && result == ""; i++ {
			if common.IsNumber(str[i]) {
				result = string(str[i])
			}
		}
	}
	return result
}

func Part1() {
	lines := common.FileToSlice("./day1/input.txt")
	total := 0
	for idx, currStr := range lines {
		foundCode := searchNearestNumeric(currStr, false) + searchNearestNumeric(currStr, true)
		value, err := strconv.Atoi(foundCode)
		if err != nil {
			fmt.Printf("Error at line %d", idx)
			panic(err)
		}
		total += value
	}
	fmt.Printf("The result is %d\n", total)
}
