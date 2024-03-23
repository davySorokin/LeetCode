class Solution:
    def jump(self, nums: List[int]) -> int:
        jumps = 0
        current_end = 0
        current_farthest = 0
        
        for i in range(len(nums) - 1):  # We do not need to jump from the last index
            current_farthest = max(current_farthest, i + nums[i])
            # If we have come to the end of the current jump, increase the jump count
            if i == current_end:
                jumps += 1
                current_end = current_farthest
        return jumps
