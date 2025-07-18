package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	// Sort the input to handle duplicates
	sort.Ints(nums)
	var res [][]int

	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		// Add the current subset to the result
		res = append(res, append([]int{}, path...))

		// Start exploring subsets from the next index
		for i := start; i < len(nums); i++ {
			// Skip duplicates
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			// Include the current element and move to the next index
			backtrack(i+1, append(path, nums[i]))
		}
	}

	// Start backtracking from index 0
	backtrack(0, []int{})
	return res
}
