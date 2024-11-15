import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Solution {
    public List<List<Integer>> findDifference(int[] nums1, int[] nums2) {
        // Create two sets to store unique elements from each array
        Set<Integer> set1 = new HashSet<>();
        Set<Integer> set2 = new HashSet<>();

        // Populate the sets with elements from nums1 and nums2
        for (int num : nums1) {
            set1.add(num);
        }
        for (int num : nums2) {
            set2.add(num);
        }

        // Create lists to store unique elements from each set
        List<Integer> uniqueInNums1 = new ArrayList<>();
        List<Integer> uniqueInNums2 = new ArrayList<>();

        // Find elements in set1 that are not in set2
        for (int num : set1) {
            if (!set2.contains(num)) {
                uniqueInNums1.add(num);
            }
        }

        // Find elements in set2 that are not in set1
        for (int num : set2) {
            if (!set1.contains(num)) {
                uniqueInNums2.add(num);
            }
        }

        // Add the lists to the answer list
        List<List<Integer>> answer = new ArrayList<>();
        answer.add(uniqueInNums1);
        answer.add(uniqueInNums2);

        return answer;
    }
}
