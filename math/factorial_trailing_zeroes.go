func trailingZeroes(n int) int {
    count := 0
    
    // Keep dividing n by 5 to count how many multiples of 5, 25, 125, etc. are there
    // Each multiple of 5 contributes at least one trailing zero in the factorial
    for n > 0 {
        n /= 5             // Divide n by 5 to get the count of numbers divisible by 5
        count += n         // Add the count of these multiples of 5 to the total
    }
    
    return count           // Return the total number of trailing zeroes
}
