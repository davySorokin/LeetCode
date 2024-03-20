import java.util.ArrayList;
import java.util.List;

class Solution {
    public int[][] insert(int[][] intervals, int[] newInterval) {
        List<int[]> result = new ArrayList<>();
        int i = 0;
        
        // Add all intervals ending before newInterval starts
        while (i < intervals.length && intervals[i][1] < newInterval[0]) {
            result.add(intervals[i]);
            i++;
        }
        
        // Merge all overlapping intervals to one that includes newInterval
        while (i < intervals.length && intervals[i][0] <= newInterval[1]) {
            newInterval = new int[] { 
                Math.min(newInterval[0], intervals[i][0]), 
                Math.max(newInterval[1], intervals[i][1])
            };
            i++;
        }
        result.add(newInterval); // Add the merged interval
        
        // Add the rest of the intervals
        while (i < intervals.length) {
            result.add(intervals[i]);
            i++;
        }
        
        // Convert list to array before returning
        return result.toArray(new int[result.size()][2]);
    }
}
