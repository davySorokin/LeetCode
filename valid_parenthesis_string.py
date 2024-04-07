class Solution:
    def checkValidString(self, s: str) -> bool:
        left_balance = right_balance = 0
        
        for c in s:
            if c == '(' or c == '*':
                left_balance += 1
            else:
                left_balance -= 1
            
            if left_balance < 0:  # More closing parentheses than opening
                return False
        
        if left_balance == 0:
            return True  # Early exit if we already know it's valid
        
        for c in reversed(s):
            if c == ')' or c == '*':
                right_balance += 1
            else:
                right_balance -= 1
            
            if right_balance < 0:  # More opening parentheses than closing when checking in reverse
                return False
        
        return True  # If neither check failed, it's valid

# usage:
# sol = Solution()
# print(sol.checkValidString("(*)"))  # Output: True
# print(sol.checkValidString("(*))"))  # Output: True
