import java.util.HashMap;
import java.util.Map;
import java.util.List;
import java.util.ArrayList;

class Solution {
    public boolean isSubsequence(String s, String t) {
        // Create a map from characters to their indices in t
        Map<Character, List<Integer>> charIndices = new HashMap<>();
        for (int i = 0; i < t.length(); i++) {
            char c = t.charAt(i);
            charIndices.putIfAbsent(c, new ArrayList<>());
            charIndices.get(c).add(i);
        }
        
        // Current index in t where we need to find the character
        int currentIndex = 0;
        for (char c : s.toCharArray()) {
            // If t doesn't have one of the characters of s, return false
            if (!charIndices.containsKey(c)) return false;
            
            // Binary search for the next occurrence of c not before currentIndex
            List<Integer> indices = charIndices.get(c);
            int i = binarySearch(indices, currentIndex);
            if (i == indices.size()) return false;  // No valid index found
            
            // Update currentIndex to the next index after the found one
            currentIndex = indices.get(i) + 1;
        }
        
        return true;  // All characters found in sequence
    }

    // Helper method to perform binary search
    // Finds the index of the smallest number in list that is >= to target
    private int binarySearch(List<Integer> list, int target) {
        int left = 0, right = list.size() - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (list.get(mid) < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return left;
    }
}
