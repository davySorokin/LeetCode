func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return cost[0]
    }

    dp := make([]int, n)
    dp[0] = cost[0]
    dp[1] = cost[1]

    for i := 2; i < n; i++ {
        dp[i] = cost[i] + min(dp[i-1], dp[i-2])
    }

    return min(dp[n-1], dp[n-2])
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
