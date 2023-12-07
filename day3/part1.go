package day3

import (
	c "aoc2023/common"
	"fmt"
	sc "strconv"
	s "strings"
)

func parseSchematic() [][]string {
	lines := c.FileToSlice("./day3/input.txt")
	schematic := [][]string{}
	for _, line := range lines {
		chars := s.Split(line, "")
		schematic = append(schematic, chars)
	}
	return schematic
}

func searchForNextSymbol(schematic [][]string, lineOffset int, charOffset int) (int, int) {
	lineIdx := -1
	charIdx := -1
	searching := true

	for currLineIdx := lineOffset; currLineIdx < len(schematic) && searching; currLineIdx++ {
		if currLineIdx == lineOffset && charOffset == len(schematic[currLineIdx]) {
			currLineIdx += 1
		}
		for currCharIdx := 0; currCharIdx < len(schematic[currLineIdx]) && searching; currCharIdx++ {
			if currLineIdx == lineOffset && currCharIdx == 0 {
				currCharIdx = charOffset + 1
			}
			currElement := schematic[currLineIdx][currCharIdx]
			isNumber := c.IsStringNumber(currElement)
			isPeriod := currElement == "."
			if !isNumber && !isPeriod {
				lineIdx = currLineIdx
				charIdx = currCharIdx
				searching = false
			}
		}
	}
	return lineIdx, charIdx
}

func searchSymbols(schematic [][]string) [][]int {
	lineOffset := 0
	charOffset := 0
	searching := true
	positions := [][]int{}
	for (lineOffset < len(schematic)-1) && (charOffset < len(schematic[0])-1) && searching {
		lineIdx, charIdx := searchForNextSymbol(schematic, lineOffset, charOffset)
		if lineIdx == -1 || charIdx == -1 {
			searching = false
		} else {
			lineOffset = lineIdx
			charOffset = charIdx
			currPosition := []int{lineIdx, charIdx}
			positions = append(positions, currPosition)
		}
	}
	return positions
}

func searchNumbersAbovePosition(schematic [][]string, linePos int, charPos int) []int {
	numbers := []int{}
	if linePos > 0 {
		middleTopNum := schematic[linePos-1][charPos]
		if c.IsStringNumber(middleTopNum) {
			// Concat all the numbers from top middle all the way to the right
			for currCharPos := charPos + 1; currCharPos < len(schematic[0]) && c.IsStringNumber(schematic[linePos-1][currCharPos]); currCharPos++ {
				middleTopNum = middleTopNum + schematic[linePos-1][currCharPos]
			}
			// Concat all the numbers from top middle all the way to the left
			for currCharPos := charPos - 1; currCharPos > -1 && c.IsStringNumber(schematic[linePos-1][currCharPos]); currCharPos-- {
				middleTopNum = schematic[linePos-1][currCharPos] + middleTopNum
			}
			fmt.Printf("middleTopNum: %s\n", middleTopNum)
			val, err := sc.Atoi(middleTopNum)
			c.Check(err)
			numbers = append(numbers, val)
		} else {
			if charPos+1 < len(schematic[0])-1 { // Check if out of range
				topRightNum := schematic[linePos-1][charPos+1]
				if c.IsStringNumber(topRightNum) {
					for currCharPos := charPos + 2; currCharPos < len(schematic[0]) && c.IsStringNumber(schematic[linePos-1][currCharPos]); currCharPos++ {
						topRightNum = topRightNum + schematic[linePos-1][currCharPos]
					}
					fmt.Printf("topRightNum: %s\n", topRightNum)
					val, err := sc.Atoi(topRightNum)
					c.Check(err)
					numbers = append(numbers, val)
				}
			}
			if charPos-1 > -1 { // Check if out of range
				topLeftNum := schematic[linePos-1][charPos-1]
				if c.IsStringNumber(topLeftNum) {
					for currCharPos := charPos - 2; currCharPos > -1 && c.IsStringNumber(schematic[linePos-1][currCharPos]); currCharPos-- {
						topLeftNum = schematic[linePos-1][currCharPos] + topLeftNum
					}
					fmt.Printf("topLeftNum: %s\n", topLeftNum)
					val, err := sc.Atoi(topLeftNum)
					c.Check(err)
					numbers = append(numbers, val)
				}
			}
		}
	}
	return numbers
}

func searchNumbersLeftPosition(schematic [][]string, linePos int, charPos int) []int {
	numbers := []int{}
	if charPos-1 > -1 { // Check if out of range
		leftNum := schematic[linePos-1][charPos-1]
		if c.IsStringNumber(leftNum) {
			for currCharPos := charPos - 2; currCharPos > -1 && c.IsStringNumber(schematic[linePos-1][currCharPos]); currCharPos-- {
				leftNum = schematic[linePos-1][currCharPos] + leftNum
			}
			fmt.Printf("leftNum: %s\n", leftNum)
			val, err := sc.Atoi(leftNum)
			c.Check(err)
			numbers = append(numbers, val)
		}
	}
	return numbers
}

func searchNumbersRightPosition(schematic [][]string, linePos int, charPos int) []int {
	numbers := []int{}
	if charPos+1 < len(schematic[0])-1 { // Check if out of range
		rightNum := schematic[linePos][charPos+1]
		if c.IsStringNumber(rightNum) {
			for currCharPos := charPos + 2; currCharPos < len(schematic[0]) && c.IsStringNumber(schematic[linePos][currCharPos]); currCharPos++ {
				rightNum = rightNum + schematic[linePos][currCharPos]
			}
			fmt.Printf("rightNum: %s\n", rightNum)
			val, err := sc.Atoi(rightNum)
			c.Check(err)
			numbers = append(numbers, val)
		}
	}
	return numbers
}

func searchNumbersBelowPosition(schematic [][]string, linePos int, charPos int) []int {
	numbers := []int{}
	if linePos < len(schematic)-1 {
		middleBottomNum := schematic[linePos+1][charPos]
		if c.IsStringNumber(middleBottomNum) {
			// Concat all the numbers from Bottom middle all the way to the right
			for currCharPos := charPos + 1; currCharPos < len(schematic[0]) && c.IsStringNumber(schematic[linePos+1][currCharPos]); currCharPos++ {
				middleBottomNum = middleBottomNum + schematic[linePos+1][currCharPos]
			}
			// Concat all the numbers from Bottom middle all the way to the left
			for currCharPos := charPos - 1; currCharPos > -1 && c.IsStringNumber(schematic[linePos+1][currCharPos]); currCharPos-- {
				middleBottomNum = schematic[linePos+1][currCharPos] + middleBottomNum
			}
			fmt.Printf("middleBottomNum: %s\n", middleBottomNum)
			val, err := sc.Atoi(middleBottomNum)
			c.Check(err)
			numbers = append(numbers, val)
		} else {
			if charPos+1 < len(schematic[0])-1 { // Check if out of range
				bottomRightNum := schematic[linePos+1][charPos+1]
				if c.IsStringNumber(bottomRightNum) {
					for currCharPos := charPos + 2; currCharPos < len(schematic[0]) && c.IsStringNumber(schematic[linePos+1][currCharPos]); currCharPos++ {
						bottomRightNum = bottomRightNum + schematic[linePos+1][currCharPos]
					}
					fmt.Printf("bottomRightNum: %s\n", bottomRightNum)
					val, err := sc.Atoi(bottomRightNum)
					c.Check(err)
					numbers = append(numbers, val)
				}
			}
			if charPos-1 > -1 { // Check if out of range
				bottomLeftNum := schematic[linePos+1][charPos-1]
				if c.IsStringNumber(bottomLeftNum) {
					for currCharPos := charPos - 2; currCharPos > -1 && c.IsStringNumber(schematic[linePos+1][currCharPos]); currCharPos-- {
						bottomLeftNum = schematic[linePos+1][currCharPos] + bottomLeftNum
					}
					fmt.Printf("bottomLeftNum: %s\n", bottomLeftNum)
					val, err := sc.Atoi(bottomLeftNum)
					c.Check(err)
					numbers = append(numbers, val)
				}
			}
		}
	}
	return numbers
}

func searchAroundPositionsForNumbers(schematic [][]string, positions [][]int) []int {
	numbers := []int{}
	for _, position := range positions {
		linePos := position[0]
		charPos := position[1]
		fmt.Printf("linePos: %d, charPos: %d\n", linePos, charPos)
		numbersOnTop := searchNumbersAbovePosition(schematic, linePos, charPos)
		numbersOnLeft := searchNumbersLeftPosition(schematic, linePos, charPos)
		numbersOnRight := searchNumbersRightPosition(schematic, linePos, charPos)
		numbersOnBottom := searchNumbersBelowPosition(schematic, linePos, charPos)
		numbers = append(numbers, numbersOnTop...)
		numbers = append(numbers, numbersOnLeft...)
		numbers = append(numbers, numbersOnRight...)
		numbers = append(numbers, numbersOnBottom...)
	}
	return numbers
}

func v1() {
	schematic := parseSchematic()
	// The schematic DS works as: schematic[line][character]
	symbolPositions := searchSymbols(schematic)
	numbers := searchAroundPositionsForNumbers(schematic, symbolPositions)
	total := c.SumArray(numbers)
	fmt.Println(total)
}

func Part1() {
	v1()
}
