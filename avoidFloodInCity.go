package main

import (
	"fmt"
	"sort"
)

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	full := make(map[int]int)
	dryDays := []int{}

	for i, lake := range rains {
		if lake == 0 {
			dryDays = append(dryDays, i)
			ans[i] = 1 // placeholder
		} else {
			if prev, ok := full[lake]; ok {
				j := sort.Search(len(dryDays), func(x int) bool { return dryDays[x] > prev })
				if j == len(dryDays) {
					return []int{}
				}
				ans[dryDays[j]] = lake
				dryDays = append(dryDays[:j], dryDays[j+1:]...)
			}
			full[lake] = i
			ans[i] = -1
		}
	}
	return ans
}

/*
func main() {
	fmt.Println(avoidFlood([]int{1, 2, 3, 4}))        // [-1,-1,-1,-1]
	fmt.Println(avoidFlood([]int{1, 2, 0, 0, 2, 1}))  // [-1,-1,2,1,-1,-1] or similar
	fmt.Println(avoidFlood([]int{1, 2, 0, 1, 2}))     // []
}
*/
