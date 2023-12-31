package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"fmt"
)

func main() {

	currDay := 3

	switch currDay {

	case 1:
		fmt.Println("Day 1 Part 1 solution: ")
		day1.Part1()
		fmt.Println("Day 1 Part 2 solution: ")
		day1.Part2()

	case 2:
		fmt.Println("Day 2 Part 1 solution: ")
		day2.Part1()
		fmt.Println("Day 2 Part 2 solution: ")
		day2.Part2()

	case 3:
		fmt.Println("Day 3 Part 1 solution: ")
		day3.Part1()
		fmt.Println("Day 3 Part 2 solution: ")
		day3.Part2()

	default:
		fmt.Println("Hello World!")
	}
}
