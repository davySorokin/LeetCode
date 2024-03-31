# runtime - 39ms
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        # Initialize count array for all ASCII characters
        count = [0] * 256  # ASCII character set size

        # Increment count for each character in s
        for char in s:
            count[ord(char)] += 1

        # Decrement count for each character in t
        for char in t:
            count[ord(char)] -= 1
            # If count for any character is negative, s and t are not anagrams
            if count[ord(char)] < 0:
                return False

        # If all counts are zero, s and t are anagrams
        return True

# Example usage:
sol = Solution()
print(sol.isAnagram("anagram", "nagaram"))  # Output: True
print(sol.isAnagram("rat", "car"))  # Output: False



# runtime - 61ms
#class Solution:
#    def isAnagram(self, s: str, t: str) -> bool:
#        # An anagram must have the same length strings.
#        if len(s) != len(t):
#            return False
#        
#        # Count characters in both strings and compare.
#        countS, countT = {}, {}
#        for i in range(len(s)):
#            countS[s[i]] = 1 + countS.get(s[i], 0)
#            countT[t[i]] = 1 + countT.get(t[i], 0)
#        
#        return countS == countT
