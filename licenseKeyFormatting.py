class Solution:
    def licenseKeyFormatting(self, s: str, k: int) -> str:
        # Remove all dashes and convert to uppercase
        s = s.replace("-", "").upper()
        
        # Determine the length of the first group
        first_group_length = len(s) % k or k
        
        # Initialise result with the first group
        result = [s[:first_group_length]]
        
        # Append the remaining groups
        for i in range(first_group_length, len(s), k):
            result.append(s[i:i+k])
        
        # Join groups with dashes
        return "-".join(result)

solution = Solution()
print(solution.licenseKeyFormatting("5F3Z-2e-9-w", 4))  # Output: "5F3Z-2E9W"
print(solution.licenseKeyFormatting("2-5g-3-J", 2))  # Output: "2-5G-3J"
