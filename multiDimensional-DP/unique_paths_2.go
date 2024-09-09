func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // Get the dimensions of the grid
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])

    // If the starting cell or ending cell is an obstacle, return 0
    if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
        return 0
    }

    // Initialize a DP array to store the number of paths
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    // The starting point has 1 path (itself)
    dp[0][0] = 1

    // Fill in the DP table
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // Skip the starting point and obstacle cells
            if i == 0 && j == 0 {
                continue
            }
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
            } else {
                // The number of paths to this cell is the sum of paths
                // from the top cell and left cell
                if i > 0 {
                    dp[i][j] += dp[i-1][j]
                }
                if j > 0 {
                    dp[i][j] += dp[i][j-1]
                }
            }
        }
    }

    // The bottom-right corner has the total number of unique paths
    return dp[m-1][n-1]
}
