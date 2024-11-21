package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := "1"

	for i := 2; i <= n; i++ {
		cur := ""
		count := 1

		for j := 1; j < len(prev); j++ {
			if prev[j] == prev[j-1] {
				count++
			} else {
				cur += strconv.Itoa(count) + string(prev[j-1])
				count = 1
			}
		}

		cur += strconv.Itoa(count) + string(prev[len(prev)-1])
		prev = cur
	}

	return prev
}
