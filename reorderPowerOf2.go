import "sort"
import "strconv"

func reorderedPowerOf2(n int) bool {
    // Generate all powers of 2 that are <= 10^9
    powers_of_2 := make(map[string]bool)
    i := 1
    for i <= 1_000_000_000 {
        powers_of_2[sortedDigits(i)] = true
        i *= 2
    }
    
    // Check if the sorted digits of n match any of the powers of 2
    return powers_of_2[sortedDigits(n)]
}

func sortedDigits(n int) string {
    digits := []rune(strconv.Itoa(n))
    sort.Slice(digits, func(i, j int) bool {
        return digits[i] < digits[j]
    })
    return string(digits)
}
