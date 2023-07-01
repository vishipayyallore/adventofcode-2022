package utilities

import (
	"fmt"
	"strconv"
)

func SolveTopThreeProblem(lines []string) int {
	calories := 0
	sumCalories := 0
	topThreeCalories := make([]int, 3)

	for _, value := range lines {
		// Print value and type. Only for debug
		// fmt.Printf("Value: %v || type:%T \n", value, value)

		c, err := strconv.Atoi(value) // Cast string to int
		if err != nil {
			c = 0
		}

		sumCalories += c

		if value == "" {
			// fmt.Println("sumCalories: ", sumCalories)

			if sumCalories > topThreeCalories[0] { // Sum Calories is greater than the first element of the array
				topThreeCalories[2] = topThreeCalories[1]
				topThreeCalories[1] = topThreeCalories[0]
				topThreeCalories[0] = sumCalories
			} else if sumCalories > topThreeCalories[1] { // Sum Calories is greater than the second element of the array
				topThreeCalories[2] = topThreeCalories[1]
				topThreeCalories[1] = sumCalories
			} else if sumCalories > topThreeCalories[2] { // Sum Calories is greater than the third element of the array
				topThreeCalories[2] = sumCalories
			}

			sumCalories = 0
		}
	}

	fmt.Printf("%v %v %v", topThreeCalories[0], topThreeCalories[1], topThreeCalories[2])

	calories = sum(topThreeCalories)

	return calories
}

func sum(arr []int) int {
	sum := 0

	for _, valueInt := range arr {
		sum += valueInt
	}

	return sum
}
