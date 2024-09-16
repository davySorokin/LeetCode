func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)

    // Create a DP table to store the minimum number of operations
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // Initialize the DP table
    for i := 0; i <= m; i++ {
        dp[i][0] = i // If word2 is empty, we need to delete all characters from word1
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j // If word1 is empty, we need to insert all characters from word2
    }

    // Fill the DP table
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1] // If characters match, no additional operations needed
            } else {
                // Consider all operations: insert, delete, replace
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
        }
    }

    return dp[m][n]
}

// Helper function to return the minimum of three integers
func min(a, b, c int) int {
    if a < b && a < c {
        return a
    }
    if b < c {
        return b
    }
    return c
}

/*
Follow-up and Constraints:

1. The function calculates the minimum number of operations (insert, delete, replace) required to convert word1 into word2. The time complexity is O(m * n), where m is the length of word1 and n is the length of word2.
2. Constraints:
   - The strings are limited to lowercase English letters.
   - The maximum length of word1 and word2 is 500, making the dynamic programming solution efficient enough.
   
   Space Optimization:
   The above implementation uses a 2D DP table with space complexity O(m * n). However, we can reduce the space complexity to O(n) by keeping only the current and previous rows of the DP table, similar to other DP space optimizations.
   
   Example of optimized space solution:

    func minDistanceOptimized(word1 string, word2 string) int {
        m, n := len(word1), len(word2)
        dp := make([]int, n+1)

        // Initialize the first row
        for j := 0; j <= n; j++ {
            dp[j] = j
        }

        // Fill the DP array
        for i := 1; i <= m; i++ {
            prev := dp[0] // previous value for dp[i-1][j-1]
            dp[0] = i     // update dp[i][0]
            for j := 1; j <= n; j++ {
                temp := dp[j] // store the current dp[i-1][j] for the next iteration
                if word1[i-1] == word2[j-1] {
                    dp[j] = prev
                } else {
                    dp[j] = min(dp[j], dp[j-1], prev) + 1
                }
                prev = temp // update previous for the next iteration
            }
        }

        return dp[n]
    }

   The optimized version reduces the space complexity to O(n), which is more efficient when handling large inputs.
*/
