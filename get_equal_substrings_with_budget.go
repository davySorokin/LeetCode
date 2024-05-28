func equalSubstring(s string, t string, maxCost int) int {
    maxLen := 0
    currentCost := 0
    start := 0

    // Calculate the differences and use a sliding window to find the maximum length
    for end := 0; end < len(s); end++ {
        // Calculate the cost to change s[end] to t[end]
        cost := abs(int(s[end]) - int(t[end]))
        currentCost += cost

        // If the current cost exceeds the maximum cost, shrink the window from the start
        for currentCost > maxCost {
            currentCost -= abs(int(s[start]) - int(t[start]))
            start++
        }

        // Update the maximum length of the substring that can be modified within the cost
        maxLen = max(maxLen, end-start+1)
    }

    return maxLen
}

// Helper functions to handle maximum and absolute value
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
