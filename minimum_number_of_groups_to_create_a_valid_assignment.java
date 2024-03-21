import java.util.*;

class Solution {
    public int minGroupsForValidAssignment(int[] balls) {
        HashMap<Integer, Integer> freq = new HashMap<>();
        int maxFreq = 0;

        // Calculate the frequency of each number.
        for (int ball : balls) {
            freq.put(ball, freq.getOrDefault(ball, 0) + 1);
            maxFreq = Math.max(maxFreq, freq.get(ball));
        }

        int minGroups = Integer.MAX_VALUE;

        // Iterate over each possible group size from 1 to the maximum frequency.
        for (int x = 1; x <= maxFreq; x++) {
            int groups = 0;
            for (int f : freq.values()) {
                // Calculate the number of groups of size x and x + 1 needed.
                int a = f / (x + 1);
                int b = f % (x + 1);
                if (b == 0) {
                    // We can form groups all of size x+1.
                    groups += a;
                } else if (x - b <= a) {
                    // We can form groups of size x and x+1.
                    groups += a + 1;
                } else {
                    // It's impossible to form groups with this size, skip to next.
                    groups = Integer.MAX_VALUE;
                    break;
                }
            }
            // Update the minimum groups if this is the lowest count we've found.
            minGroups = Math.min(minGroups, groups);
        }

        return minGroups;
    }
}
