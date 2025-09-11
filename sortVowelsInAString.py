class Solution:
    def sortVowels(self, s: str) -> str:
        vowels = set('aeiouAEIOU')
        s_list = list(s)

        # Extract and sort the vowels
        vowel_chars = sorted([c for c in s if c in vowels])
        
        # Replace vowels in-place in order
        vowel_index = 0
        for i in range(len(s_list)):
            if s_list[i] in vowels:
                s_list[i] = vowel_chars[vowel_index]
                vowel_index += 1

        return ''.join(s_list)

# usage/checking output
sol = Solution()
print(sol.sortVowels("lEetcOde"))  # Output: "lEOtcede"
print(sol.sortVowels("lYmpH"))     # Output: "lYmpH"
