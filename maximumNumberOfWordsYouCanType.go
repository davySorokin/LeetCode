func canBeTypedWords(text string, brokenLetters string) int {
    words := strings.Split(text, " ")
    broken := make(map[rune]bool)
    for _, ch := range brokenLetters {
        broken[ch] = true
    }

    count := 0
    for _, word := range words {
        canType := true
        for _, ch := range word {
            if broken[ch] {
                canType = false
                break
            }
        }
        if canType {
            count++
        }
    }
    return count
}
