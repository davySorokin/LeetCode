class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""
        
        # Start by assuming the whole first string is the common prefix
        common_prefix = strs[0]
        
        for s in strs[1:]:
            # Check each string in the list against the current common prefix
            i = 0
            while i < len(s) and i < len(common_prefix) and s[i] == common_prefix[i]:
                i += 1
            # Update the common prefix to the matched portion
            common_prefix = common_prefix[:i]
            
            # If at any time the common prefix becomes empty, we can return immediately
            if not common_prefix:
                return ""
        
        return common_prefix

# Example usage:
# sol = Solution()
# print(sol.longestCommonPrefix(["flower","flow","flight"])) # Output: "fl"
# print(sol.longestCommonPrefix(["dog","racecar","car"])) # Output: ""
