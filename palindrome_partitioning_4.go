func checkPartitioning(s string) bool {
    n := len(s)
    
    // Precompute a DP table to check if substring s[i:j+1] is a palindrome
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }

    // Fill the DP table
    for length := 1; length <= n; length++ {
        for i := 0; i+length-1 < n; i++ {
            j := i + length - 1
            if s[i] == s[j] {
                if length <= 2 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            }
        }
    }

    // Check for valid partitioning into three palindromic parts
    for i := 1; i < n-1; i++ {
        for j := i; j < n-1; j++ {
            if dp[0][i-1] && dp[i][j] && dp[j+1][n-1] {
                return true
            }
        }
    }

    return false
}

/*
This solution uses dynamic programming to precompute all possible palindromic substrings, which ensures the solution runs in O(n^2) time complexity, 
where n is the length of the string. The space complexity is also O(n^2) due to the DP table. 
This approach efficiently handles the given constraint of 3 <= s.length <= 2000.
*/
