class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        results = []

        def backtrack(curr_str="", open_n=0, close_n=0):
            # Base case: if both opening and closing parentheses are balanced, add the combination
            if open_n == n and close_n == n:
                results.append(curr_str)
                return

            # Add opening parenthesis if possible
            if open_n < n:
                backtrack(curr_str + "(", open_n + 1, close_n)

            # Add closing parenthesis if valid (more closing than opening is not allowed)
            if close_n < open_n:
                backtrack(curr_str + ")", open_n, close_n + 1)

        backtrack()
        return results
