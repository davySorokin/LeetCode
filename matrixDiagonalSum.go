func diagonalSum(mat [][]int) int {
    n := len(mat)
    sum := 0

    for i := 0; i < n; i++ {
        // Add primary
        sum += mat[i][i]
        
        // Add secondary
        sum += mat[i][n-i-1]
    }
    
    // If n is odd, subtract the middle element (since it was added twice)
    if n % 2 == 1 {
        sum -= mat[n/2][n/2]
    }

    return sum
}
