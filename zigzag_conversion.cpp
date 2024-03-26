#include <vector>
#include <string>

class Solution {
public:
    string convert(string s, int numRows) {
        // If there is only one row or the string is shorter than the number of rows,
        // then the string itself is already in zigzag form.
        if (numRows == 1 || numRows >= s.length()) {
            return s;
        }

        vector<string> rows(min(numRows, int(s.length())));
        int curRow = 0;
        bool goingDown = false;

        // Traverse the string and place characters in the appropriate row.
        for (char c : s) {
            rows[curRow] += c;
            if (curRow == 0 || curRow == numRows - 1) {
                goingDown = !goingDown; // Change direction at the first or last row
            }
            curRow += goingDown ? 1 : -1; // Move to the next row
        }

        // Concatenate all rows to form the final string.
        string result;
        for (string row : rows) {
            result += row;
        }

        return result;
    }
};
