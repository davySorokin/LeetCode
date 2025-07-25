class Solution:
    def isScramble(self, s1: str, s2: str) -> bool:
        if len(s1) != len(s2):
            return False
        
        memo = {}

        def dfs(s1, s2):
            if s1 == s2:
                return True

            if len(s1) == 1:
                return s1 == s2

            if (s1, s2) in memo:
                return memo[(s1, s2)]

            n = len(s1)
            for i in range(1, n):
                # Case 1: Don't swap
                if dfs(s1[:i], s2[:i]) and dfs(s1[i:], s2[i:]):
                    memo[(s1, s2)] = True
                    return True
                # Case 2: Swap
                if dfs(s1[:i], s2[n-i:]) and dfs(s1[i:], s2[:n-i]):
                    memo[(s1, s2)] = True
                    return True

            memo[(s1, s2)] = False
            return False

        return dfs(s1, s2)
