func isScramble(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }

    memo := make(map[string]bool)
    
    var dfs func(s1, s2 string) bool
    dfs = func(s1, s2 string) bool {
        if s1 == s2 {
            return true
        }

        if len(s1) == 1 {
            return s1 == s2
        }

        key := s1 + "-" + s2
        if result, exists := memo[key]; exists {
            return result
        }

        n := len(s1)
        for i := 1; i < n; i++ {
            // Case 1: Don't swap
            if dfs(s1[:i], s2[:i]) && dfs(s1[i:], s2[i:]) {
                memo[key] = true
                return true
            }
            // Case 2: Swap
            if dfs(s1[:i], s2[n-i:]) && dfs(s1[i:], s2[:n-i]) {
                memo[key] = true
                return true
            }
        }

        memo[key] = false
        return false
    }

    return dfs(s1, s2)
}
