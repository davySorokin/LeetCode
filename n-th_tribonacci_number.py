class Solution:
    def tribonacci(self, n: int) -> int:
        # Initialize a memoization dictionary to store previously calculated values
        memo = {0: 0, 1: 1, 2: 1}

        # Define a recursive helper function to calculate the Tribonacci number
        def helper(n):
            if n in memo:
                return memo[n]
            memo[n] = helper(n-1) + helper(n-2) + helper(n-3)
            return memo[n]

        # Call the helper function to get the nth Tribonacci number
        return helper(n)
