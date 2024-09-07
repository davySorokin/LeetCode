func plusOne(digits []int) []int {
    n := len(digits)
    
    // Traverse the array from right to left
    for i := n - 1; i >= 0; i-- {
        digits[i]++
        
        if digits[i] < 10 {
            return digits
        }
        
        // If the current digit becomes 10, we set it to 0 and continue
        digits[i] = 0
    }
    
    // If we have gone through all digits and they are all 0, we need to add 1 at the beginning
    return append([]int{1}, digits...)
}
