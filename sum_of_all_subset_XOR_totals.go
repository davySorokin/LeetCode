package main

import "fmt"

func subsetXORSum(nums []int) int {
    return recSubsetXORSum(nums, 0, 0)
}

// Helper function to use recursion for finding XOR of all subsets
func recSubsetXORSum(nums []int, index int, currentXOR int) int {
    if index == len(nums) {
        return currentXOR
    }
    // Include the current number in the subset
    include := recSubsetXORSum(nums, index+1, currentXOR^nums[index])
    // Exclude the current number from the subset
    exclude := recSubsetXORSum(nums, index+1, currentXOR)
    return include + exclude
}
