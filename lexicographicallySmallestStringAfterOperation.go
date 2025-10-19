package main

import (
	"fmt"
)

func findLexSmallestString(s string, a int, b int) string {
	seen := make(map[string]bool)
	queue := []string{s}
	res := s

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if seen[cur] {
			continue
		}
		seen[cur] = true

		if cur < res {
			res = cur
		}

		// Operation 1: add a to all digits at odd indices
		arr := []byte(cur)
		for i := 1; i < len(arr); i += 2 {
			arr[i] = byte((int(arr[i]-'0')+a)%10 + '0')
		}
		added := string(arr)
		if !seen[added] {
			queue = append(queue, added)
		}

		// Operation 2: rotate right by b positions
		n := len(cur)
		rotated := cur[n-b:] + cur[:n-b]
		if !seen[rotated] {
			queue = append(queue, rotated)
		}
	}

	return res
}
/*
func main() {
	fmt.Println(findLexSmallestString("5525", 9, 2)) // Output: 2050
	fmt.Println(findLexSmallestString("74", 5, 1))   // Output: 24
	fmt.Println(findLexSmallestString("0011", 4, 2)) // Output: 0011
}
*/
