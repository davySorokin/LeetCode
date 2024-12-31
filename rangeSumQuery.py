class NumArray:

    def __init__(self, nums: list[int]):
        self.prefix_sums = [0] * (len(nums) + 1)
        for i in range(len(nums)):
            self.prefix_sums[i + 1] = self.prefix_sums[i] + nums[i]

    def sumRange(self, left: int, right: int) -> int:
        return self.prefix_sums[right + 1] - self.prefix_sums[left]


# Example usage: <-- Not used in Leetcode.
nums = [-2, 0, 3, -5, 2, -1]
numArray = NumArray(nums)
print(numArray.sumRange(0, 2))  # Output: 1
print(numArray.sumRange(2, 5))  # Output: -1
print(numArray.sumRange(0, 5))  # Output: -3
