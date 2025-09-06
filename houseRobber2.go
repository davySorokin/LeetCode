func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    return max(robLinear(nums[:len(nums)-1]), robLinear(nums[1:]))
}

func robLinear(houses []int) int {
    prev, curr := 0, 0
    for _, money := range houses {
        prev, curr = curr, max(curr, prev+money)
    }
    return curr
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
