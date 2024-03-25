import java.util.*;

class Solution {
    public List<Integer> findDuplicates(int[] nums) {
        List<Integer> duplicates = new ArrayList<>();

        for (int num : nums) {
            // Use the absolute value of the current number as an index
            int index = Math.abs(num) - 1;
            
            // If the value at this index is negative, it's a duplicate
            if (nums[index] < 0) {
                duplicates.add(Math.abs(num));
            }
            
            // Negate the value at this index to mark it as visited
            nums[index] = -nums[index];
        }

        return duplicates;
    }
}
