func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    original := x
    reversed := 0
    
    for x > 0 {
        lastDigit := x % 10
        reversed = reversed * 10 + lastDigit
        x /= 10
    }
    
    return original == reversed
}
