func areOccurrencesEqual(s string) bool {
    freq := make(map[rune]int)
    
    for _, ch := range s {
        freq[ch]++
    }

    // Check: if all frequencies are the same
    var occurrence int
    for _, count := range freq {
        if occurrence == 0 {
            occurrence = count
        } else if occurrence != count {
            return false
        }
    }
    return true
}
