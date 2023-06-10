package main

import (
	"adventofcode/day1/utilities"
	"fmt"
)

func main() {
	fmt.Println("Advent of Code - Day 1 - Go - 2022")

	fileNameWithPath := "./data/input.txt" // "./data/input.txt" OR "./data/sample.txt"

	lines := utilities.GetInput(fileNameWithPath)
	// fmt.Println(lines)  // Print all lines. Only for debug

	calories := utilities.SolveProblem(lines)

	fmt.Println("Calories: ", calories)
}
