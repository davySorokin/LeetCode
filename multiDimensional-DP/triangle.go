func minimumTotal(triangle [][]int) int {
    // Start from the second last row of the triangle and move upwards.
    for i := len(triangle) - 2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            // The value at each point is updated to the minimum sum of the path
            // by considering both the adjacent positions in the next row.
            triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
        }
    }
    // The top of the triangle will now contain the minimum path sum.
    return triangle[0][0]
}

// Helper function to find the minimum of two integers.
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
