func nextPermutation(nums []int) {
    n := len(nums)
    if n < 2 {
        return
    }

    // Step 1: Find the first decreasing element from the end
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }

    if i >= 0 {
        // Step 2: Find the element larger than nums[i] to swap with
        j := n - 1
        for nums[j] <= nums[i] {
            j--
        }

        // Step 3: Swap
        nums[i], nums[j] = nums[j], nums[i]
    }

    // Step 4: Reverse the portion after i
    reverse(nums, i+1, n-1)
}

// Helper function to reverse the array from start to end
func reverse(nums []int, start, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
