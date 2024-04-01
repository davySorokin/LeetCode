#include <vector>
#include <algorithm>

class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        if (intervals.empty()) return {};

        // Sort the intervals based on the start time
        sort(intervals.begin(), intervals.end(),
             [](const vector<int>& a, const vector<int>& b) {
                 return a[0] < b[0];
             });

        vector<vector<int>> merged;
        for (const vector<int>& interval : intervals) {
            // If the list of merged intervals is empty or if the current
            // interval does not overlap with the previous, simply append it.
            if (merged.empty() || merged.back()[1] < interval[0]) {
                merged.push_back(interval);
            } else {
                // Otherwise, there is overlap, so we merge the current and previous
                // intervals by updating the end of the previous interval if needed.
                merged.back()[1] = max(merged.back()[1], interval[1]);
            }
        }

        return merged;
    }
};
