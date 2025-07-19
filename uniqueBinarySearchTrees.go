func numTrees(n int) int {
    // dp[i] will store the number of unique BSTs that can be formed with 'i' nodes
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1 // Base case: 1 way to form BST with 0 or 1 node
    
    // Calculate the number of unique BSTs for all values from 2 to n
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            dp[i] += dp[j-1] * dp[i-j] // Left * Right subtrees
        }
    }
    
    return dp[n]
}
