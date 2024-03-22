class Solution {
    public int splitArray(int[] nums, int k) {
        int max = 0; // The max number in the array
        long sum = 0; // Sum of all numbers in the array
        for (int num : nums) {
            max = Math.max(num, max);
            sum += num;
        }
        
        // Binary search between the max number and the sum of all numbers
        long lo = max;
        long hi = sum;
        while (lo < hi) {
            long mid = lo + (hi - lo) / 2;
            if (valid(mid, nums, k)) {
                hi = mid;
            } else {
                lo = mid + 1;
            }
        }
        
        return (int) lo;
    }
    
    private boolean valid(long target, int[] nums, int k) {
        int count = 1;
        long total = 0;
        for (int num : nums) {
            total += num;
            if (total > target) {
                total = num;
                count++;
                if (count > k) {
                    return false;
                }
            }
        }
        return true;
    }
}
