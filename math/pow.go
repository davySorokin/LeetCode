func myPow(x float64, n int) float64 {
    // Handle the case where n is negative
    if n < 0 {
        x = 1 / x
        n = -n
    }
    
    result := 1.0
    currentProduct := x
    
    // Exponentiation by squaring (iterative approach)
    for n > 0 {
        // If n is odd, multiply the result by the current product
        if n % 2 == 1 {
            result *= currentProduct
        }
        
        // Square the base and reduce n by half
        currentProduct *= currentProduct
        n /= 2
    }
    
    return result
}
