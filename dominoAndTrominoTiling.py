class Solution:
    def numTilings(self, n: int) -> int:
        MOD = 10**9 + 7
        if n == 1:
            return 1
        dp = [0] * (n + 1)
        dp2 = [0] * (n + 1)
        dp[0], dp[1] = 1, 1
        dp2[0] = 0

        for i in range(2, n + 1):
            dp[i] = (dp[i - 1] + dp[i - 2] + 2 * dp2[i - 1]) % MOD
            dp2[i] = (dp2[i - 1] + dp[i - 2]) % MOD

        return dp[n]
