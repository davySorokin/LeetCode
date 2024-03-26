class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        # Return the index of the first occurrence of needle in haystack
        return haystack.find(needle)

# Example usage:
# sol = Solution()
# print(sol.strStr("sadbuttsad", "sad"))  # Output: 0
# print(sol.strStr("leetcode", "leeto"))  # Output: -1
