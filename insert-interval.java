class Solution {
    public int[][] insert(int[][] intervals, int[] newInterval) {
        List<int[]> result = new ArrayList<>();

        for (int[] interval : intervals) {
            // If the current interval ends before the new interval starts, add it to the result.
            if (interval[1] < newInterval[0]) {
                result.add(interval);
            // If the current interval starts after the new interval ends, add the new interval to the result.
            // Then update the new interval to the current one since we are done with the new interval.
            } else if (interval[0] > newInterval[1]) {
                result.add(newInterval);
                newInterval = interval;
            // Otherwise, there is an overlap, so we merge the current interval with the new interval.
            } else {
                newInterval[0] = Math.min(newInterval[0], interval[0]);
                newInterval[1] = Math.max(newInterval[1], interval[1]);
            }
        }

        // Add the last interval which could be the new interval or a merge of intervals.
        result.add(newInterval);

        return result.toArray(new int[result.size()][]);
    }
}
