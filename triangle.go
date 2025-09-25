package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	// dp holds best costs from current row down
	dp := make([]int, n)
	// initialize with last row
	for i := 0; i < n; i++ {
		dp[i] = triangle[n-1][i]
	}
	// process upwards
	for r := n - 2; r >= 0; r-- {
		for c := 0; c <= r; c++ {
			if dp[c] < dp[c+1] {
				dp[c] = triangle[r][c] + dp[c]
			} else {
				dp[c] = triangle[r][c] + dp[c+1]
			}
		}
	}
	return dp[0]
}
