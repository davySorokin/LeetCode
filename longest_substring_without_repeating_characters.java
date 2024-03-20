import java.util.HashMap;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        HashMap<Character, Integer> map = new HashMap<>();
        int maxLength = 0;
        int left = 0; // Left side of the sliding window
        
        for (int right = 0; right < s.length(); right++) {
            char currentChar = s.charAt(right);
            if (map.containsKey(currentChar)) {
                // Move the left side of the window to the right of the first occurrence of currentChar
                left = Math.max(left, map.get(currentChar) + 1);
            }
            map.put(currentChar, right); // Update the last seen index of the currentChar
            maxLength = Math.max(maxLength, right - left + 1); // Calculate the max length so far
        }
        
        return maxLength;
    }
}
