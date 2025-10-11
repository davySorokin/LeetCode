import (
    "sort"
)

func maximumTotalDamage(power []int) int64 {
    count := map[int]int{}
    for _, v := range power {
        count[v]++
    }

    unique := make([]int, 0, len(count))
    for k := range count {
        unique = append(unique, k)
    }
    sort.Ints(unique)

    n := len(unique)
    dp := make([][4]int64, n+1)

    for i := 1; i <= n; i++ {
        val := unique[i-1]
        total := int64(val * count[val])
        // binary search for last index allowed
        prev := sort.SearchInts(unique, val-2)
        for j := 0; j < 4; j++ {
            dp[i][j] = dp[i-1][j]
            if j > 0 && dp[i-1][j-1] > dp[i][j] {
                dp[i][j] = dp[i-1][j-1]
            }
            if dp[prev][j]+total > dp[i][j] {
                dp[i][j] = dp[prev][j] + total
            }
        }
    }
    ans := int64(0)
    for j := 0; j < 4; j++ {
        if dp[n][j] > ans {
            ans = dp[n][j]
        }
    }
    return ans
}
