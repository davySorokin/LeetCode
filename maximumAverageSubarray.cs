public class Solution {
    public double FindMaxAverage(int[] nums, int k) {
        double maxSum = 0, windowSum = 0;

        for (int i = 0; i < k; i++) {
            windowSum += nums[i];
        }
        maxSum = windowSum;

        for (int i = k; i < nums.Length; i++) {
            windowSum += nums[i] - nums[i - k];
            if (windowSum > maxSum) {
                maxSum = windowSum;
            }
        }

        return maxSum / k;
    }
}
