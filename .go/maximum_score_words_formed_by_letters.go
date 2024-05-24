func maxScoreWords(words []string, letters []byte, score []int) int {
    // Letter frequency count
    letterCount := make([]int, 26)
    for _, letter := range letters {
        letterCount[letter-'a']++
    }

    // Calculate score for each word
    var calculateWordScore func(word string) int
    calculateWordScore = func(word string) int {
        tempCount := make([]int, 26)
        wordScore := 0
        for _, ch := range word {
            idx := ch - 'a'
            tempCount[idx]++
            if tempCount[idx] > letterCount[idx] { // More letters needed than available
                return 0
            }
            wordScore += score[idx]
        }
        return wordScore
    }

    // Recursive function to try all subsets of words
    var maxScore int
    var dfs func(index int, currentScore int)
    dfs = func(index int, currentScore int) {
        if index == len(words) {
            if currentScore > maxScore {
                maxScore = currentScore
            }
            return
        }

        // Do not include the current word
        dfs(index+1, currentScore)

        // Try to include the current word
        wordScore := calculateWordScore(words[index])
        if wordScore > 0 { // Only proceed if the word can be formed
            // Update letter counts
            for _, ch := range words[index] {
                letterCount[ch-'a']--
            }

            // Include the word in the subset
            dfs(index+1, currentScore+wordScore)

            // Backtrack, restore letter counts
            for _, ch := range words[index] {
                letterCount[ch-'a']++
            }
        }
    }

    // Start DFS from the first word
    dfs(0, 0)
    return maxScore
}
