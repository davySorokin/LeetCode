func maxFreqSum(s string) int {
    freq := make(map[rune]int)
    vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
    
    for _, ch := range s {
        freq[ch]++
    }
    
    maxVowel := 0
    maxConsonant := 0
    
    for ch, count := range freq {
        if vowels[ch] {
            if count > maxVowel {
                maxVowel = count
            }
        } else {
            if count > maxConsonant {
                maxConsonant = count
            }
        }
    }
    
    return maxVowel + maxConsonant
}
