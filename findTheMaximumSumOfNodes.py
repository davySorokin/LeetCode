class Solution:
    def maximumValueSum(self, nums: List[int], k: int, edges: List[List[int]]) -> int:
        base_sum = sum(nums)
        gains = [((num ^ k) - num) for num in nums]
        gains.sort(reverse=True)
        
        total_gain = 0
        count = 0
        for i in range(0, len(gains) - 1, 2):
            pair_gain = gains[i] + gains[i + 1]
            if pair_gain > 0:
                total_gain += pair_gain
                count += 2
            else:
                break
        
        return base_sum + total_gain
