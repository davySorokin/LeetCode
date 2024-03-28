#include <vector>
#include <string>
#include <unordered_map>
#include <climits>

class Solution {
public:
    std::string minWindow(std::string s, std::string t) {
        if (s.empty() || t.empty()) {
            return "";
        }

        // Create a frequency map for the characters in t
        std::unordered_map<char, int> t_freq;
        for (char c : t) {
            t_freq[c]++;
        }

        // Initialize the pointers for the sliding window
        int left = 0, right = 0;
        int required = t_freq.size();
        int formed = 0;
        std::unordered_map<char, int> window_counts;

        // Pair of the form (window length, (left, right))
        std::pair<int, std::pair<int, int>> ans = {INT_MAX, {0, 0}};

        while (right < s.size()) {
            // Add one character from the right to the window
            char c = s[right];
            window_counts[c]++;
            
            // If the frequency of the current character added equals to the
            // frequency of the character in t, then increment the formed count
            if (t_freq.count(c) && window_counts[c] == t_freq[c]) {
                formed++;
            }

            // Try and contract the window till the point where it ceases to be 'desirable'
            while (left <= right && formed == required) {
                c = s[left];
                // Save the smallest window until now
                if (right - left + 1 < ans.first) {
                    ans = {right - left + 1, {left, right}};
                }

                // The character at the position pointed by the `left` pointer is no longer a part of the window
                window_counts[c]--;
                if (t_freq.count(c) && window_counts[c] < t_freq[c]) {
                    formed--;
                }

                // Move the left pointer ahead, this would help to look for a new window
                left++;
            }

            // Keep expanding the window once we are done contracting
            right++;
        }

        return ans.first == INT_MAX ? "" : s.substr(ans.second.first, ans.first);
    }
};
