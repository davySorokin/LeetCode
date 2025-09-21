func countMaxOrSubsets(nums []int) int {
    n := len(nums)
    maxOr := 0
    for _, v := range nums {
        maxOr |= v
    }

    // tail[i] = OR of nums[i:]
    tail := make([]int, n+1)
    for i := n - 1; i >= 0; i-- {
        tail[i] = tail[i+1] | nums[i]
    }

    var dfs func(i, cur int) int
    dfs = func(i, cur int) int {
        if i == n {
            if cur == maxOr {
                return 1
            }
            return 0
        }
        // prune if impossible to reach maxOr
        if (cur|tail[i]) < maxOr {
            return 0
        }
        res := dfs(i+1, cur)             // skip
        res += dfs(i+1, cur|nums[i])     // take
        return res
    }

    return dfs(0, 0)
}
