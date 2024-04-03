#include <vector>
#include <string>

class Solution {
public:
    bool dfs(vector<vector<char>>& board, int i, int j, string& word, int index) {
        if (index == word.size()) {
            return true;
        }
        if (i < 0 || i >= board.size() || j < 0 || j >= board[0].size() || board[i][j] != word[index]) {
            return false;
        }
        char temp = board[i][j];
        board[i][j] = '#'; // Mark as visited
        bool found = dfs(board, i + 1, j, word, index + 1) ||
                     dfs(board, i - 1, j, word, index + 1) ||
                     dfs(board, i, j + 1, word, index + 1) ||
                     dfs(board, i, j - 1, word, index + 1);
        board[i][j] = temp; // Revert the change for backtracking
        return found;
    }

    bool exist(vector<vector<char>>& board, string word) {
        for (int i = 0; i < board.size(); ++i) {
            for (int j = 0; j < board[i].size(); ++j) {
                if (board[i][j] == word[0] && dfs(board, i, j, word, 0)) {
                    return true;
                }
            }
        }
        return false;
    }
};
