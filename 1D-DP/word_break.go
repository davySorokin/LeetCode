func wordBreak(s string, wordDict []string) bool {
    // Constraint 1: s.length <= 300
    // This will be handled as the algorithm works by breaking down the string
    // into substrings and checking against wordDict. We use dynamic programming (DP)
    // to avoid redundant checks and keep the complexity low.
    
    // Constraint 2: wordDict.length <= 1000
    // A hash set is used to efficiently check if a substring is in the dictionary,
    // which allows constant time lookups, and can handle up to 1000 dictionary words.
    wordSet := make(map[string]bool)
    for _, word := range wordDict {
        wordSet[word] = true
    }

    // DP array to keep track if a substring up to index i can be segmented
    dp := make([]bool, len(s)+1)
    dp[0] = true // Base case: an empty string is always "segmentable"
    
    // Constraint 3: wordDict[i].length <= 20
    // By iterating from each possible position in the string (up to 300),
    // we only consider substrings of length up to 20.
    for i := 1; i <= len(s); i++ {
        for j := i - 1; j >= 0 && i-j <= 20; j-- { // Check substrings of length <= 20
            if dp[j] && wordSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }

    return dp[len(s)] // Returns true if the entire string can be segmented
}
