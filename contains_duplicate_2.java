import java.util.HashSet;
import java.util.Set;

class Solution {
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        Set<Integer> set = new HashSet<>();

        for (int i = 0; i < nums.length; i++) {
            if (i > k) {
                // Remove the number that's now out of the k-distance window
                set.remove(nums[i - k - 1]);
            }

            // If the number is already in the set, we found a duplicate within k distance
            if (!set.add(nums[i])) {
                return true;
            }
        }

        return false;
    }
}
