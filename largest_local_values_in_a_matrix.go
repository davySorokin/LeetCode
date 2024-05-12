package main

import (
    "fmt"
)

// largestLocal returns a new matrix where each element is the maximum of the
// 3x3 submatrix centered around the corresponding element in the input matrix.
func largestLocal(grid [][]int) [][]int {
    n := len(grid)    // Assuming grid is always square
    m := n - 2        // Size of the new grid
    result := make([][]int, m)
    for i := range result {
        result[i] = make([]int, m)
    }

    for i := 1; i <= n-2; i++ {
        for j := 1; j <= n-2; j++ {
            maxVal := grid[i-1][j-1] // Start by assuming the top-left corner of the 3x3 submatrix is the max
            for di := -1; di <= 1; di++ {
                for dj := -1; dj <= 1; dj++ {
                    if grid[i+di][j+dj] > maxVal {
                        maxVal = grid[i+di][j+dj]
                    }
                }
            }
            result[i-1][j-1] = maxVal
        }
    }

    return result
}
