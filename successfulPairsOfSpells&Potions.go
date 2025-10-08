package main

func successfulPairsBrute(spells []int, potions []int, success int64) []int {
	n := len(spells)
	ans := make([]int, n)
	for i, s := range spells {
		var cnt int
		ss := int64(s)
		for _, p := range potions {
			if ss*int64(p) >= success {
				cnt++
			}
		}
		ans[i] = cnt
	}
	return ans
}
