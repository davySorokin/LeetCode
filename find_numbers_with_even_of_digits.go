package main

import (
	"fmt"
)

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		if numDigits(num)%2 == 0 {
			count++
		}
	}
	return count
}

func numDigits(n int) int {
	digits := 0
	for n > 0 {
		n /= 10
		digits++
	}
	return digits
}
