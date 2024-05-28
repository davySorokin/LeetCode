package main

import (
    "fmt"
    "sort"
    "math"
)

func threeSumClosest(nums []int, target int) int {
    // Sort the array first
    sort.Ints(nums)
    closest := nums[0] + nums[1] + nums[2]  // Initial guess using the first three elements
    
    // Iterate over each element and use two pointers to find the closest sum
    for i := 0; i < len(nums)-2; i++ {
        left, right := i+1, len(nums)-1
        // Use two pointers to explore possible sums
        for left < right {
            currentSum := nums[i] + nums[left] + nums[right]
            // Update closest sum if the current sum is closer to the target
            if math.Abs(float64(target-currentSum)) < math.Abs(float64(target-closest)) {
                closest = currentSum
            }
            // Move the pointers based on comparison
            if currentSum < target {
                left++
            } else if currentSum > target {
                right--
            } else {
                // If the current sum is exactly the target, it's the closest possible
                return currentSum
            }
        }
    }
    return closest
}
