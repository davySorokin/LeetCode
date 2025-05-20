func isZeroArray(nums []int, queries [][]int) bool {
    n := len(nums)
    count := make([]int, n+1)

    for _, q := range queries {
        l, r := q[0], q[1]
        count[l]++
        if r+1 < len(count) {
            count[r+1]--
        }
    }

    for i := 1; i < n; i++ {
        count[i] += count[i-1]
    }

    for i := 0; i < n; i++ {
        if nums[i] > count[i] {
            return false
        }
    }

    return true
}
