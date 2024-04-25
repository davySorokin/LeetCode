class Solution:
    def longestIdealString(self, s: str, k: int) -> int:
        n = len(s)
        
        # Handle empty string case
        if n == 0:
            return 0

        # Initialize dp table with all zeros (size 27 for lowercase letters)
        dp = [0] * 27

        # Base case: last character is a subsequence of length 1
        dp[ord(s[-1]) - ord('a')] = 1

        # Iterate string right to left
        for i in range(n - 2, -1, -1):
            # Get index in dp table for current character
            char_idx = ord(s[i]) - ord('a')

            # Define range of possible previous characters (considering k)
            min_prev_idx = max(0, char_idx - k)
            max_prev_idx = min(26, char_idx + k)

            # Find maximum value in dp within the range
            max_prev_val = 0
            for prev_idx in range(min_prev_idx, max_prev_idx + 1):
                max_prev_val = max(max_prev_val, dp[prev_idx])

            # Update dp for current character
            dp[char_idx] = max_prev_val + 1

        # Return the maximum value in dp (longest subsequence)
        return max(dp)
