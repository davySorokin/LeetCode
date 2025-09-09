class Solution:
    def peopleAwareOfSecret(self, n: int, delay: int, forget: int) -> int:
        MOD = 10**9 + 7
        dp = [0] * (n + 1)  # dp[i] = new people who learn secret on day i
        dp[1] = 1
        total = 0

        for day in range(1, n + 1):
            sharers = dp[day]
            for share_day in range(day + delay, min(n + 1, day + forget)):
                dp[share_day] = (dp[share_day] + sharers) % MOD

        for day in range(n - forget + 1, n + 1):
            total = (total + dp[day]) % MOD

        return total
