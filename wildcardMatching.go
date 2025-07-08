func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([]bool, n+1)
    prev := make([]bool, n+1)
    prev[0] = true

    for j := 1; j <= n; j++ {
        if p[j-1] == '*' {
            prev[j] = prev[j-1]
        }
    }

    for i := 1; i <= m; i++ {
        dp = make([]bool, n+1)
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                dp[j] = dp[j-1] || prev[j]
            } else if p[j-1] == '?' || s[i-1] == p[j-1] {
                dp[j] = prev[j-1]
            }
        }
        prev = dp
    }
    return prev[n]
}
