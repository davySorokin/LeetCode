package main

import (
	"fmt"
	"sort"
)

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	diags := map[int][]int{}

	// Collect values by diagonal (i - j)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d := i - j
			diags[d] = append(diags[d], grid[i][j])
		}
	}

	// Sort diagonals depending on region
	for d, vals := range diags {
		if d >= 0 {
			sort.Sort(sort.Reverse(sort.IntSlice(vals))) // non-increasing
		} else {
			sort.Ints(vals) // non-decreasing
		}
		diags[d] = vals
	}

	// Write back to grid
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d := i - j
			grid[i][j] = diags[d][0]
			diags[d] = diags[d][1:]
		}
	}

	return grid
}

/*
// testing
func main() {
	grid := [][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}
	fmt.Println(sortMatrix(grid))
}
*/
