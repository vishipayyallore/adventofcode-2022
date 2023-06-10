package utilities

import "strconv"

func SolveProblem(lines []string) int {
	calories := 0
	sumCalories := 0

	for _, value := range lines {
		// Print value and type. Only for debug
		// fmt.Printf("Value: %v || type:%T \n", value, value)

		c, err := strconv.Atoi(value) // Cast string to int
		if err != nil {
			c = 0
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
