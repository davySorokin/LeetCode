func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // Initialize the maximum, minimum, and result with the first element
    maxProd, minProd, result := nums[0], nums[0], nums[0]

    // Iterate over the array starting from the second element
    for i := 1; i < len(nums); i++ {
        current := nums[i]
        
        // If the current number is negative, swap max and min
        if current < 0 {
            maxProd, minProd = minProd, maxProd
        }
        
        // Update max and min products considering the current number
        maxProd = max(current, maxProd*current)
        minProd = min(current, minProd*current)
        
        // Update the result with the maximum product found so far
        result = max(result, maxProd)
    }

    return result
}

// Helper functions for min and max
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/*
This solution runs in O(n) time, which is efficient given the constraint that nums can have up to 2 * 10^4 elements. 
The space complexity is O(1), as we are only using a few extra variables (maxProd, minProd, result) to track the current state. 
The constraints ensure that the product of any subarray will fit within a 32-bit integer, so no overflow concerns.
*/
