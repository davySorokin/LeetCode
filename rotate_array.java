class Solution {
    public void rotate(int[] nums, int k) {
        if (nums == null || nums.length < 2) {
            return;
        }
        
        int n = nums.length;
        k = k % n; // In case the number of rotations is more than the length of array
        reverse(nums, 0, n - 1); // Reverse the whole array
        reverse(nums, 0, k - 1); // Reverse first k elements
        reverse(nums, k, n - 1); // Reverse last n-k elements
    }

    private void reverse(int[] nums, int start, int end) {
        while (start < end) {
            int temp = nums[start];
            nums[start] = nums[end];
            nums[end] = temp;
            start++;
            end--;
        }
    }
}
