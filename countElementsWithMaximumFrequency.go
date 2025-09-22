func maxFrequencyElements(nums []int) int {
    freq := make(map[int]int)
    maxFreq := 0

    // Count frequencies
    for _, num := range nums {
        freq[num]++
        if freq[num] > maxFreq {
            maxFreq = freq[num]
        }
    }

    total := 0
    for _, count := range freq {
        if count == maxFreq {
            total += count
        }
    }

    return total
}
