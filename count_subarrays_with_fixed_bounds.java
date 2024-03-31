class Solution {
    public long countSubarrays(int[] nums, int minK, int maxK) {
        long count = 0;
        int lastMin = -1, lastMax = -1, lastOut = -1;
        
        for (int i = 0; i < nums.length; i++) {
            // Update the last seen positions for minK, maxK and out of range elements.
            if (nums[i] < minK || nums[i] > maxK) {
                lastOut = i;
            }
            if (nums[i] == minK) {
                lastMin = i;
            }
            if (nums[i] == maxK) {
                lastMax = i;
            }
            
            // The number of valid subarrays ending at the current position is determined
            // by the furthest lastMin and lastMax, but starting after the last out of range element.
            if (lastMin != -1 && lastMax != -1) {
                count += Math.max(0, Math.min(lastMin, lastMax) - lastOut);
            }
        }
        
        return count;
    }
}
