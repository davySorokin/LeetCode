class Solution:
    def colorTheGrid(self, m: int, n: int) -> int:
        from itertools import product
        from collections import defaultdict

        MOD = 10**9 + 7

        # Generate all valid columns (tuples of length m)
        def is_valid(col):
            return all(col[i] != col[i+1] for i in range(len(col)-1))

        colors = [0, 1, 2]  # 0: red, 1: green, 2: blue
        valid_cols = [col for col in product(colors, repeat=m) if is_valid(col)]

        # Build transitions between valid columns
        transitions = defaultdict(list)
        for a in valid_cols:
            for b in valid_cols:
                if all(x != y for x, y in zip(a, b)):
                    transitions[a].append(b)

        # DP initialisation
        dp = {col: 1 for col in valid_cols}

        # Fill DP table for each column from 2 to n
        for _ in range(n - 1):
            new_dp = defaultdict(int)
            for col in dp:
                for nxt in transitions[col]:
                    new_dp[nxt] = (new_dp[nxt] + dp[col]) % MOD
            dp = new_dp

        return sum(dp.values()) % MOD
