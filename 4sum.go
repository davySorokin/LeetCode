package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	// Sort the input array
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}

	// Iterate through the array
	for i := 0; i < n-3; i++ {
		// Avoid duplicates for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// Avoid duplicates for the second number
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j + 1, n - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					// Avoid duplicates for the third number
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					// Avoid duplicates for the fourth number
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}
