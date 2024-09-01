import "math"

func mctFromLeafValues(arr []int) int {
    n := len(arr)
    dp := make([][]int, n)
    maxLeaf := make([][]int, n)
    
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        maxLeaf[i] = make([]int, n)
        maxLeaf[i][i] = arr[i]
        for j := i + 1; j < n; j++ {
            maxLeaf[i][j] = max(maxLeaf[i][j-1], arr[j])
        }
    }

    for length := 2; length <= n; length++ {
        for i := 0; i <= n-length; i++ {
            j := i + length - 1
            dp[i][j] = math.MaxInt32
            for k := i; k < j; k++ {
                dp[i][j] = min(dp[i][j], dp[i][k] + dp[k+1][j] + maxLeaf[i][k] * maxLeaf[k+1][j])
            }
        }
    }

    return dp[0][n-1]
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
