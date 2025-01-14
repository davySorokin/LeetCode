# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num: int) -> int:

class Solution:
    def guessNumber(self, n: int) -> int:
        left, right = 1, n  # Initialise the binary search boundaries
        
        while left <= right:
            mid = (left + right) // 2  # Find the middle point
            result = guess(mid)  # Call the guess API
            
            if result == 0:  # Found the number
                return mid
            elif result == -1:  # Guess is higher, move to the left
                right = mid - 1
            else:  # Guess is lower, move to the right
                left = mid + 1
