func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    dp := make([]int, n)
    maxVal := energy[n-1]

    for i := n - 1; i >= 0; i-- {
        dp[i] = energy[i]
        if i+k < n {
            dp[i] += dp[i+k]
        }
        if dp[i] > maxVal {
            maxVal = dp[i]
        }
    }
    return maxVal
}
