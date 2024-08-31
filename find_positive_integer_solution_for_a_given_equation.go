package main

import (
	"fmt"
)

func findSolution(customFunction func(int, int) int, z int) [][]int {
	result := [][]int{}
	x, y := 1, 1000

	for x <= 1000 && y >= 1 {
		value := customFunction(x, y)
		if value == z {
			result = append(result, []int{x, y})
			x++
			y--
		} else if value < z {
			x++
		} else {
			y--
		}
	}

	return result
}
