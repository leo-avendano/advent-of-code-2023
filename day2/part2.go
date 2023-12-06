package day2

import (
	c "aoc2023/common"
	"fmt"
)

func Part2() {
	games := c.FileToSlice("./day2/input.txt")
	result := 0
	for _, game := range games {
		_, rounds := extractGame(game)
		validGame := true
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for i := 0; i < len(rounds) && validGame; i++ {
			red, green, blue := extractColors(rounds[i])
			if red > maxRed {
				maxRed = red
			}
			if green > maxGreen {
				maxGreen = green
			}
			if blue > maxBlue {
				maxBlue = blue
			}
		}
		result += (maxRed * maxGreen * maxBlue)
	}
	fmt.Printf("Result is: %d \n", result)
}
