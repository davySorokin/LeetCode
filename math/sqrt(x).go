func mySqrt(x int) int {
    // Edge case: the square root of 0 is 0
    if x == 0 {
        return 0
    }

    // Binary search between 1 and x
    left, right := 1, x
    
    // This loop will explore all integers between 1 and x (Hint 1),
    // using the sorted nature of integers to narrow the search (Hint 2).
    for left <= right {
        mid := left + (right - left) / 2 // Avoid potential overflow
        
        // If mid^2 equals x, return mid as the exact square root
        if mid*mid == x {
            return mid
        } else if mid*mid < x {
            // If mid squared is less than x, explore the right half
            left = mid + 1
        } else {
            // If mid squared is more than x, explore the left half
            right = mid - 1
        }
    }
    
    // After the loop, 'right' will be the largest integer whose square is <= x
    return right
}
