package main

import (
	"adventofcode/day1/utilities"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code - Day 1 - Go - 2022")

	// fileNameWithPath := "./data/sample.txt"
	fileNameWithPath := "./data/input.txt"

	lines := utilities.GetInput(fileNameWithPath)

	// Print all lines. Only for debug
	// fmt.Println(lines)

	calories := solveProblem(lines)

	fmt.Println("Calories: ", calories)
}

func solveProblem(lines []string) int {
	calories := 0
	sumCalories := 0

	for _, value := range lines {
		// Print value and type. Only for debug
		// fmt.Printf("Value: %v || type:%T \n", value, value)

		c, err := strconv.Atoi(value) // Cast string to int
		if err != nil {
			print("Error occurred at newDeckFromFile() - ", err)
			os.Exit(1)
		}

		sumCalories += c

		if value == "" {
			if sumCalories > calories {
				calories = sumCalories
			}
			sumCalories = 0
		}
	}
	return calories
}
