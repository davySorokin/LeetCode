import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var backtrack func(start int, target int, path []int)

	backtrack = func(start int, target int, path []int) {
		if target == 0 {
			comb := make([]int, len(path))
			copy(comb, path)
			res = append(res, comb)
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue // skip duplicates
			}
			if candidates[i] > target {
				break
			}
			backtrack(i+1, target-candidates[i], append(path, candidates[i]))
		}
	}
	backtrack(0, target, []int{})
	return res
}
