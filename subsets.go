func subsets(nums []int) [][]int {
    var res [][]int
    var curr []int

    var dfs func(int)
    dfs = func(start int) {
        // Make a copy of the current subset and add to result
        subset := make([]int, len(curr))
        copy(subset, curr)
        res = append(res, subset)

        // Explore further elements
        for i := start; i < len(nums); i++ {
            curr = append(curr, nums[i])     // choose
            dfs(i + 1)                        // explore
            curr = curr[:len(curr)-1]        // un-choose (backtrack)
        }
    }

    dfs(0)
    return res
}
