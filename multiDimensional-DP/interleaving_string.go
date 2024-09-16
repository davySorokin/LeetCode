func isInterleave(s1 string, s2 string, s3 string) bool {
    // If lengths don't match, s3 cannot be an interleaving of s1 and s2
    if len(s1) + len(s2) != len(s3) {
        return false
    }

    // DP table to store the state of interleaving
    dp := make([][]bool, len(s1)+1)
    for i := range dp {
        dp[i] = make([]bool, len(s2)+1)
    }

    dp[0][0] = true // Empty s1 and s2 interleave to form empty s3

    // Fill the DP table
    for i := 0; i <= len(s1); i++ {
        for j := 0; j <= len(s2); j++ {
            if i > 0 {
                dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
            }
            if j > 0 {
                dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
            }
        }
    }

    return dp[len(s1)][len(s2)]
}

/*
Follow-up and Constraints:

1. The function checks if s3 is formed by interleaving s1 and s2. The interleaving condition means characters are taken alternately from s1 and s2 while maintaining their internal order.
2. Constraints: 
   - If the lengths of s1, s2, and s3 don't sum correctly (len(s1) + len(s2) != len(s3)), return false immediately.
   - The code uses dynamic programming, which ensures the time complexity is O(m * n), where m = len(s1) and n = len(s2).
   
   Space Optimization:
   The above implementation uses a 2D DP array with O(m * n) space complexity. However, this can be optimized to O(n) by using a 1D DP array, where we reuse the previous row's values. This is possible because each DP state only depends on the current row and the previous row.

   Here's how you can optimize the space complexity to O(n):

    func isInterleaveOptimized(s1 string, s2 string, s3 string) bool {
        if len(s1) + len(s2) != len(s3) {
            return false
        }

        dp := make([]bool, len(s2)+1)
        dp[0] = true

        for j := 1; j <= len(s2); j++ {
            dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
        }

        for i := 1; i <= len(s1); i++ {
            dp[0] = dp[0] && s1[i-1] == s3[i-1]
            for j := 1; j <= len(s2); j++ {
                dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) || (dp[j-1] && s2[j-1] == s3[i+j-1])
            }
        }

        return dp[len(s2)]
    }
    
   The optimized version reduces space complexity to O(n), where n is the length of s2.
*/
