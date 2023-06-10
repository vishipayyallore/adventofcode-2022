package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 1 - Go - 2022")

	// fileNameWithPath := "./data/sample.txt"
	fileNameWithPath := "./data/input.txt"

	lines := getInput(fileNameWithPath)

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

		c, _ := strconv.Atoi(value) // Cast string to int
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

func getInput(filename string) []string {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		print("Error occurred at newDeckFromFile() - ", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\r\n")
}
