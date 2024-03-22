class Solution {
    public int majorityElement(int[] nums) {
        // Initialize candidate and counter
        int count = 0;
        Integer candidate = null;

        for (int num : nums) {
            // If count is 0, we set the current candidate to num
            if (count == 0) {
                candidate = num;
            }
            // If the num is the candidate, we increment count, otherwise decrement
            count += (num == candidate) ? 1 : -1;
        }

        // Candidate is now the majority element
        return candidate;
    }
}
