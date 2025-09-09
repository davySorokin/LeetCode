func peopleAwareOfSecret(n int, delay int, forget int) int {
    mod := int(1e9 + 7)
    dp := make([]int, n+1)
    dp[1] = 1
    total := 0

    for day := 1; day <= n; day++ {
        sharers := dp[day]
        for shareDay := day + delay; shareDay < day + forget && shareDay <= n; shareDay++ {
            dp[shareDay] = (dp[shareDay] + sharers) % mod
        }
    }

    for day := n - forget + 1; day <= n; day++ {
        total = (total + dp[day]) % mod
    }

    return total
}
