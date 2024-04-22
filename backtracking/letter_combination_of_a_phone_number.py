from typing import List

class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        
        # Map each digit to its corresponding letters
        phone_map = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz",
        }
        
        def backtrack(index, path):
            # If the path is the same length as digits, we have a complete combination
            if len(path) == len(digits):
                combinations.append("".join(path))
                return
            # Get the letters that the current digit can represent
            possible_letters = phone_map[digits[index]]
            # Loop through these letters
            for letter in possible_letters:
                # Add the letter to our current path and move on to the next digit
                path.append(letter)
                backtrack(index + 1, path)
                # Backtrack by removing the letter before moving onto the next
                path.pop()
        
        combinations = []
        backtrack(0, [])
        return combinations
