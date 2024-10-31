import (
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	n, m := len(robot), len(factory)
	maxValue := int64(1<<53 - 1)
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, m+1)
		for j := range dp[i] {
			dp[i][j] = maxValue
		}
	}
	dp[0][0] = 0

	for j := 1; j <= m; j++ {
		pos, limit := factory[j-1][0], factory[j-1][1]
		for i := 0; i <= n; i++ {
			dp[i][j] = dp[i][j-1]
			sum := int64(0)
			for k := 1; k <= limit && i-k >= 0; k++ {
				sum += int64(abs(robot[i-k] - pos))
				dp[i][j] = min(dp[i][j], dp[i-k][j-1]+sum)
			}
		}
	}
	return dp[n][m]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
