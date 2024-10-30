func minimumMountainRemovals(nums []int) int {
    n := len(nums)
    
    // Step 1: Calculate LIS (Longest Increasing Subsequence) from left
    left := make([]int, n)
    for i := 0; i < n; i++ {
        left[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                left[i] = max(left[i], left[j] + 1)
            }
        }
    }
    
    // Step 2: Calculate LIS from right (Longest Decreasing Subsequence)
    right := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        right[i] = 1
        for j := n - 1; j > i; j-- {
            if nums[j] < nums[i] {
                right[i] = max(right[i], right[j] + 1)
            }
        }
    }
    
    // Step 3: Find maximum length of mountain array
    maxMountain := 0
    for i := 1; i < n-1; i++ {
        if left[i] > 1 && right[i] > 1 { // Must be a valid peak
            maxMountain = max(maxMountain, left[i] + right[i] - 1)
        }
    }
    
    // Step 4: Calculate the minimum removals to form mountain array
    return n - maxMountain
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
