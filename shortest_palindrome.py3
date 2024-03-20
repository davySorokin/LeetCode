class Solution:
    def shortestPalindrome(self, s: str) -> str:
        if not s:
            return s

        # Concatenate string | delimiter | reversed string
        temp = s + "|" + s[::-1]
        # Create the LPS (longest prefix suffix) array
        lps = [0] * len(temp)
        prev_lps, i = 0, 1

        while i < len(temp):
            if temp[i] == temp[prev_lps]:
                prev_lps += 1
                lps[i] = prev_lps
                i += 1
            elif prev_lps != 0:
                prev_lps = lps[prev_lps - 1]
            else:
                lps[i] = 0
                i += 1

        # Last value of lps will tell us the count of characters to be added to the front
        to_add_count = len(s) - lps[-1]
        # Add the required characters to the front to make it a palindrome
        return s[-1:-to_add_count-1:-1] + s
