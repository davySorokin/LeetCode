import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    var res [][]int
    sort.Ints(candidates)

    var backtrack func(start int, target int, path []int)
    backtrack = func(start int, target int, path []int) {
        if target == 0 {
            combo := make([]int, len(path))
            copy(combo, path)
            res = append(res, combo)
            return
        }

        for i := start; i < len(candidates); i++ {
            if i > start && candidates[i] == candidates[i-1] {
                continue // skip duplicates
            }
            if candidates[i] > target {
                break // no point in continuing
            }

            path = append(path, candidates[i])
            backtrack(i+1, target-candidates[i], path)
            path = path[:len(path)-1] // backtrack
        }
    }

    backtrack(0, target, []int{})
    return res
}
