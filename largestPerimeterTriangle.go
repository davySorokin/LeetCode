package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // sort descending
	for i := 0; i < len(nums)-2; i++ {
		a, b, c := nums[i], nums[i+1], nums[i+2]
		if b+c > a { // triangle inequality
			return a + b + c
		}
	}
	return 0
}

/*
func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))     // 5
	fmt.Println(largestPerimeter([]int{1, 2, 1, 10})) // 0
}
*/
