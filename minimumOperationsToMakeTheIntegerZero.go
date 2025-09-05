package main

import (
	"math/bits"
)

func makeTheIntegerZero(num1 int, num2 int) int {
	for k := 1; k <= 60; k++ {
		S := int64(num1) - int64(k)*int64(num2)
		if S <= 0 {
			if num2 >= 0 {
				break
			}
			continue
		}
		// popcount(S) <= k <= S
		pop := bits.OnesCount64(uint64(S))
		if pop <= k && int64(k) <= S {
			return k
		}
	}
	return -1
}
