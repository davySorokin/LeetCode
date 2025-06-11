import java.util.Arrays;

class Solution {
    public int maximumProduct(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        
        // O.1: Product of the three largest numbers
        int option1 = nums[n - 1] * nums[n - 2] * nums[n - 3];

        // O.2: Product of the two smallest numbers and the largest number
        int option2 = nums[0] * nums[1] * nums[n - 1];

        return Math.max(option1, option2);
    }
}
