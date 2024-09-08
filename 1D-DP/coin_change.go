func coinChange(coins []int, amount int) int {
    // Handle constraint: coins is a non-empty array (as per the constraint).
    // We assume an infinite supply of coins of each denomination as per the problem description.
    
    // Create a dp array where dp[i] will store the minimum number of coins required to make up amount i.
    dp := make([]int, amount+1)
    
    // Initialize the dp array. Set a large value (amount + 1) because we'll minimize it later.
    for i := 1; i <= amount; i++ {
        dp[i] = amount + 1
    }
    
    // Base case: 0 coins are required to make up amount 0.
    dp[0] = 0
    
    // Iterate through each amount from 1 to the given amount.
    for i := 1; i <= amount; i++ {
        // Try every coin denomination in coins to see if it can be used to minimize the number of coins.
        for _, coin := range coins {
            if i-coin >= 0 {
                dp[i] = min(dp[i], dp[i-coin]+1)
            }
        }
    }
    
    // If the dp[amount] is still amount+1, it means it's impossible to make up that amount with the given coins.
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

// Helper function to find the minimum of two numbers.
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
