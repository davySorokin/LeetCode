class Solution {
    public int findMaxLength(int[] nums) {
        // Map to store the earliest index where a particular count is found
        Map<Integer, Integer> countMap = new HashMap<>();
        countMap.put(0, -1); // Base case for when the count is zero
        int maxLength = 0, count = 0;

        for (int i = 0; i < nums.length; i++) {
            count += (nums[i] == 1) ? 1 : -1; // Increment count for 1, decrement for 0

            if (countMap.containsKey(count)) {
                // If the count has been seen before, a subarray with equal numbers of 0s and 1s has been found
                maxLength = Math.max(maxLength, i - countMap.get(count));
            } else {
                // Store the index for this count
                countMap.put(count, i);
            }
        }

        return maxLength;
    }
}
