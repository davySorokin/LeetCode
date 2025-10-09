class Solution:
    def minTime(self, skill: List[int], mana: List[int]) -> int:
        n, m = len(skill), len(mana)
        dp = [0] * m

        for i in range(n):
            new_dp = [0] * m
            for j in range(m):
                time = skill[i] * mana[j]
                if i == 0 and j == 0:
                    new_dp[j] = time
                elif i == 0:
                    new_dp[j] = new_dp[j-1] + time
                elif j == 0:
                    new_dp[j] = dp[j] + time
                else:
                    new_dp[j] = max(dp[j], new_dp[j-1]) + time
            dp = new_dp

        return dp[-1]
