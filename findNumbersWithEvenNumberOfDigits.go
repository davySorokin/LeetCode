package main

import {
  "fmt"
  "strconv"
  }

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			count++
		}
	}
	return count
}

// tests within
func main() {
	nums1 := []int{12, 345, 2, 6, 7896}
	fmt.Println("Output:", findNumbers(nums1)) // Output: 2
  
	nums2 := []int{555, 901, 482, 1771}
	fmt.Println("Output:", findNumbers(nums2)) // Output: 1
}
