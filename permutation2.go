import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums) // Here -> Sort to ensure we can skip duplicates
	backtrack(nums, []int{}, &res, make([]bool, len(nums)))
	return res
}

func backtrack(nums []int, path []int, res *[][]int, used []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	k-- // convert to 0-based index
    	res := ""
	
	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		used[i] = true
		backtrack(nums, append(path, nums[i]), res, used)
		used[i] = false
	}
}
