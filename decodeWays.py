class Solution:
    def numDecodings(self, s: str) -> int:
        # Edge case: if the string starts with '0', no valid decoding is possible
        if not s or s[0] == '0':
            return 0
        
        # Initialize dp array, where dp[i] is the number of ways to decode s[0:i]
        n = len(s)
        dp = [0] * (n + 1)
        dp[0] = 1  # Base case: There's one way to decode an empty string
        dp[1] = 1  # There's one way to decode the first character if it's not '0'

        for i in range(2, n + 1):
            # Check for single character decoding
            if s[i - 1] != '0':
                dp[i] += dp[i - 1]
            
            # Check for two character decoding
            if s[i - 2] == '1' or (s[i - 2] == '2' and '1' <= s[i - 1] <= '6'):
                dp[i] += dp[i - 2]

        return dp[n]
