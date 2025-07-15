class Solution {
    func search(_ nums: [Int], _ target: Int) -> Bool {
        var left = 0, right = nums.count - 1

        while left <= right {
            let mid = left + (right - left) / 2

            if nums[mid] == target {
                return true
            }

            // Handle duplicates
            if nums[left] == nums[mid] {
                left += 1
            } else if nums[left] < nums[mid] {
                // Left sorted
                if nums[left] <= target && target < nums[mid] {
                    right = mid - 1
                } else {
                    left = mid + 1
                }
            } else {
                // Right sorted
                if nums[mid] < target && target <= nums[right] {
                    left = mid + 1
                } else {
                    right = mid - 1
                }
            }
        }

        return false
    }
}
