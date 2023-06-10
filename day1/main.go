package main

import (
	"adventofcode/day1/utilities"
	"fmt"
)

func main() {
	fmt.Println("Advent of Code - Day 1 - Go - 2022")

	// fileNameWithPath := "./data/sample.txt"
	fileNameWithPath := "./data/input.txt"

	lines := utilities.GetInput(fileNameWithPath)
	// Print all lines. Only for debug
	// fmt.Println(lines)

	calories := utilities.SolveProblem(lines)

	fmt.Println("Calories: ", calories)
}
