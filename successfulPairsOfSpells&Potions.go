package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)
	res := make([]int, len(spells))

	for i, s := range spells {
		need := (success + int64(s) - 1) / int64(s) // ceil(success / s)
		idx := sort.Search(m, func(k int) bool {
			return int64(potions[k]) >= need
		})
		res[i] = m - idx
	}

	return res
}
