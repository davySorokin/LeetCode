package main

import "fmt"

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)

	for i := 0; i < n; i++ {
		left[i] = 1
		right[i] = 1
	}

	// increasing lengths ending at i
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	// increasing lengths starting at i
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		k := left[i]
		if right[i+1] < k {
			k = right[i+1]
		}
		if k > ans {
			ans = k
		}
	}
	return ans
}

/*
func main() {
	fmt.Println(maxIncreasingSubarrays([]int{2,5,7,8,9,2,3,4,3,1})) // 3
	fmt.Println(maxIncreasingSubarrays([]int{1,2,3,4,4,4,4,5,6,7})) // 2
}
*/
