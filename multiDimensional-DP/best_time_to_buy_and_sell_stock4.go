func maxProfit(k int, prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }
    
    // If k is large enough, it becomes equivalent to unlimited transactions
    if k >= n/2 {
        return maxProfitUnlimited(prices)
    }

    // DP table to store the max profit
    dp := make([][]int, k+1)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    // Fill DP table
    for t := 1; t <= k; t++ {
        maxDiff := -prices[0]  // max difference (buying and selling in the past)
        for i := 1; i < n; i++ {
            dp[t][i] = max(dp[t][i-1], prices[i]+maxDiff)  // either don't sell or sell at current price
            maxDiff = max(maxDiff, dp[t-1][i]-prices[i])  // update the max difference
        }
    }

    return dp[k][n-1]
}

// Helper function for unlimited transactions (k is large enough)
func maxProfitUnlimited(prices []int) int {
    maxProfit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            maxProfit += prices[i] - prices[i-1]  // accumulate all positive differences
        }
    }
    return maxProfit
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
