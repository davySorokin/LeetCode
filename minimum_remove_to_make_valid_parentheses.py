class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        # First pass to remove extra closing parentheses
        balance = 0
        result = []
        for char in s:
            if char == '(':
                balance += 1
            elif char == ')':
                if balance == 0:
                    continue
                balance -= 1
            result.append(char)
        
        # Second pass to remove extra opening parentheses from the end
        i = len(result) - 1
        while balance > 0 and i >= 0:
            if result[i] == '(':
                result.pop(i)
                balance -= 1
            i -= 1
            
        return ''.join(result)
