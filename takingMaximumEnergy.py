class Solution:
    def maximumEnergy(self, energy: list[int], k: int) -> int:
        n = len(energy)
        dp = [0] * n
        
        for i in range(n - 1, -1, -1):
            dp[i] = energy[i]
            if i + k < n:
                dp[i] += dp[i + k]
        
        return max(dp)
