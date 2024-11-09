package main

import (
	"fmt"
)

func maxProductPath(grid [][]int) int {
	const MOD = 1_000_000_007
	m, n := len(grid), len(grid[0])

	maxProduct := make([][]int, m)
	minProduct := make([][]int, m)
	for i := 0; i < m; i++ {
		maxProduct[i] = make([]int, n)
		minProduct[i] = make([]int, n)
	}

	maxProduct[0][0] = grid[0][0]
	minProduct[0][0] = grid[0][0]

	for j := 1; j < n; j++ {
		maxProduct[0][j] = maxProduct[0][j-1] * grid[0][j]
		minProduct[0][j] = maxProduct[0][j]
	}

	for i := 1; i < m; i++ {
		maxProduct[i][0] = maxProduct[i-1][0] * grid[i][0]
		minProduct[i][0] = maxProduct[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] >= 0 {
				maxProduct[i][j] = max(maxProduct[i-1][j], maxProduct[i][j-1]) * grid[i][j]
				minProduct[i][j] = min(minProduct[i-1][j], minProduct[i][j-1]) * grid[i][j]
			} else {
				maxProduct[i][j] = min(minProduct[i-1][j], minProduct[i][j-1]) * grid[i][j]
				minProduct[i][j] = max(maxProduct[i-1][j], maxProduct[i][j-1]) * grid[i][j]
			}
		}
	}

	// Get: result from the bottom-right corner
	result := maxProduct[m-1][n-1]
	if result < 0 {
		return -1
	}
	return result % MOD
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
