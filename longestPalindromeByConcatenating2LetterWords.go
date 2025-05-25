func longestPalindrome(words []string) int {
    count := make(map[string]int)
    for _, word := range words {
        count[word]++
    }

    length := 0
    centralUsed := false

    for word, c := range count {
        rev := string([]byte{word[1], word[0]})

        if word == rev {
            // Pair palindromic words like "gg"
            pairs := c / 2
            length += pairs * 4
            count[word] -= pairs * 2
            if !centralUsed && count[word] > 0 {
                length += 2
                centralUsed = true
            }
        } else if word < rev { // avoid double counting
            if revCount, ok := count[rev]; ok {
                pairs := min(c, revCount)
                length += pairs * 4
                count[word] -= pairs
                count[rev] -= pairs
            }
        }
    }

    return length
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
