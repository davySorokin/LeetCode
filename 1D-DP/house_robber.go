func rob(nums []int) int {
    // Edge case: if there are no houses, the max amount to rob is 0.
    if len(nums) == 0 {
        return 0
    }

    // Edge case: if there's only one house, rob it.
    if len(nums) == 1 {
        return nums[0]
    }

    // We use two variables to track the max amount we can rob:
    // prev1 will store the max amount till the previous house.
    // prev2 will store the max amount till the house before the previous house.
    prev1, prev2 := 0, 0

    // Iterate through each house in the array.
    for _, num := range nums {
        // The max amount to rob this house would either be:
        // - The amount robbed till the previous house (prev1), or
        // - The amount robbed till the house before the previous one (prev2) plus the current house's amount.
        newRob := max(prev1, prev2+num)

        // Update prev2 to be the previous house's amount and
        // prev1 to be the new max we calculated.
        prev2 = prev1
        prev1 = newRob
    }

    // At the end of the loop, prev1 holds the maximum amount we can rob.
    return prev1
}

// Helper function to calculate the maximum of two numbers.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
