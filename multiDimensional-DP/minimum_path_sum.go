func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    
    // Traverse the grid and update each cell with the minimum path sum up to that cell.
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                // Skip the starting cell as it remains the same.
                continue
            } else if i == 0 {
                // If we're in the first row, we can only come from the left.
                grid[i][j] += grid[i][j-1]
            } else if j == 0 {
                // If we're in the first column, we can only come from above.
                grid[i][j] += grid[i-1][j]
            } else {
                // Take the minimum of coming from the left or from above.
                grid[i][j] += min(grid[i-1][j], grid[i][j-1])
            }
        }
    }
    
    // The bottom-right cell now contains the minimum path sum.
    return grid[m-1][n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
