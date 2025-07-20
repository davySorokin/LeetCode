func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' { return 0 }
    dp := make([]int, len(s)+1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= len(s); i++ {
        if s[i-1] != '0' { dp[i] += dp[i-1] }
        if s[i-2:i] >= "10" && s[i-2:i] <= "26" { dp[i] += dp[i-2] }
    }
    return dp[len(s)]
}
