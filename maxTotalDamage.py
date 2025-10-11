from collections import Counter
from bisect import bisect_right

class Solution:
    def maximumTotalDamage(self, power):
        count = Counter(power)
        unique = sorted(count)
        n = len(unique)
        dp = [[0]*4 for _ in range(n+1)]  # dp[i][j] = max dmg using first i spells, skipping j unique spells

        for i in range(1, n+1):
            val = unique[i-1]
            total = val * count[val]
            # find prev index we can take (no conflict within 2 range)
            prev = bisect_right(unique, val - 3)
            for j in range(4):
                # skip this spell
                dp[i][j] = max(dp[i][j], dp[i-1][j])
                if j > 0:
                    dp[i][j] = max(dp[i][j], dp[i-1][j-1])
                # take this spell
                dp[i][j] = max(dp[i][j], dp[prev][j] + total)
        return max(dp[n])
