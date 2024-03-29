class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        # Helper function to count subarrays where the target is the max element appearing at least k times
        def count_for_max(arr, target, k):
            n = len(arr)
            ans = 0
            cnt = 0
            left = 0
            for right in range(n):
                if arr[right] == target:
                    cnt += 1
                while cnt >= k:
                    ans += (n - right)
                    if arr[left] == target:
                        cnt -= 1
                    left += 1
            return ans
        
        # Main function to find the overall count
        maxi = max(nums)
        return count_for_max(nums, maxi, k)

# Example usage
# sol = Solution()
# print(sol.countSubarrays([1,3,2,3,3], 2))  # Expected output: 6
# print(sol.countSubarrays([1,4,2,1], 3))    # Expected output: 0
