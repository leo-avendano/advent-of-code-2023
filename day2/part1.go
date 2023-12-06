package day2

import (
	c "aoc2023/common"
	"fmt"
	sc "strconv"
	s "strings"
)

func extractGame(str string) (int, []string) {
	substrings := s.Split(str, ":")
	gameId, err := sc.Atoi(substrings[0][5:])
	rounds := s.Split(substrings[1], ";")
	c.Check(err)
	return gameId, rounds
}

func extractColors(round string) (int, int, int) {
	red := 0
	green := 0
	blue := 0
	results := s.Split(round, ",")
	for _, result := range results {
		substr := s.Split(s.Trim(result, " "), " ")
		amount, err := sc.Atoi(substr[0])
		c.Check(err)
		color := substr[1]
		switch color {
		case "red":
			red = amount
		case "green":
			green = amount
		case "blue":
			blue = amount
		default:
			panic("Invalid color")
		}
	}
	return red, green, blue
}

func Part1() {
	games := c.FileToSlice("./day2/input.txt")
	result := 0
	maxRedCubes := 12
	maxGreenCubes := 13
	maxBlueCubes := 14
	for _, game := range games {
		gameId, rounds := extractGame(game)
		validGame := true
		for i := 0; i < len(rounds) && validGame; i++ {
			red, green, blue := extractColors(rounds[i])
			if red > maxRedCubes || green > maxGreenCubes || blue > maxBlueCubes {
				validGame = false
			}
		}
		if validGame {
			result += gameId
		}
	}
	fmt.Printf("Result is: %d \n", result)
}
