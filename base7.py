class Solution:
    def convertToBase7(self, num: int) -> str:
        if num == 0:
            return "0"

        negative = num < 0
        num = abs(num)
        res = []

        while num:
            res.append(str(num % 7))
            num //= 7

        if negative:
            res.append('-')

        return ''.join(res[::-1])

sol = Solution()
print(sol.convertToBase7(100))  # Output: "202"
print(sol.convertToBase7(-7))   # Output: "-10"
