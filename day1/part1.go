package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNumber(b byte) bool {
	return (b >= 48 && b <= 57)
}

func fileToSlice(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func searchNearestNumeric(str string, reverse bool) string {
	result := ""
	if reverse {
		for i := len(str) - 1; i >= 0 && result == ""; i-- {
			if isNumber(str[i]) {
				result = string(str[i])
			}
		}
	} else {
		for i := 0; i < len(str) && result == ""; i++ {
			if isNumber(str[i]) {
				result = string(str[i])
			}
		}
	}
	return result
}

func part1() {
	lines := fileToSlice("input.txt")
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
