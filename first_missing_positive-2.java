class Solution {
    public int firstMissingPositive(int[] nums) {
        int n = nums.length;

        for (int i = 0; i < n; ++i) {
            // While the current number is in the range of 1 to n, and it's not in its target position,
            // which is nums[i] - 1, swap it to its correct position.
            while (nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i]) {
                int correct = nums[i] - 1;
                int temp = nums[i];
                nums[i] = nums[correct];
                nums[correct] = temp;
            }
        }

        // After rearrangement, the first index that doesn't match the number is the answer.
        for (int i = 0; i < n; ++i) {
            if (nums[i] != i + 1) {
                return i + 1;
            }
        }

        // If all numbers are in place, the missing number is n+1.
        return n + 1;
    }
}
