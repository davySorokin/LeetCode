func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }

    rows, cols := len(matrix), len(matrix[0])
    maxSide := 0
    dp := make([][]int, rows+1)
    
    // Initialize DP array with zeros
    for i := range dp {
        dp[i] = make([]int, cols+1)
    }

    // Iterate through matrix to fill dp
    for i := 1; i <= rows; i++ {
        for j := 1; j <= cols; j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
                maxSide = max(maxSide, dp[i][j])
            }
        }
    }

    // Return area of the largest square
    return maxSide * maxSide
}

// Helper function to get the minimum of three integers
func min(a, b, c int) int {
    if a < b && a < c {
        return a
    }
    if b < c {
        return b
    }
    return c
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
