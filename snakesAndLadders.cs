public class Solution {
    public int SnakesAndLadders(int[][] board) {
        int n = board.Length;
        int[] flatten = new int[n * n + 1];
        int idx = 1;
        bool leftToRight = true;

        for (int i = n - 1; i >= 0; i--) {
            if (!leftToRight) Array.Reverse(board[i]);
            foreach (int val in board[i]) {
                flatten[idx++] = val;
            }
            leftToRight = !leftToRight;
        }

        bool[] visited = new bool[n * n + 1];
        Queue<(int pos, int moves)> q = new();
        q.Enqueue((1, 0));
        visited[1] = true;

        while (q.Count > 0) {
            var (curr, moves) = q.Dequeue();
            for (int i = 1; i <= 6; i++) {
                int next = curr + i;
                if (next > n * n) break;
                if (flatten[next] != -1) next = flatten[next];
                if (next == n * n) return moves + 1;
                if (!visited[next]) {
                    visited[next] = true;
                    q.Enqueue((next, moves + 1));
                }
            }
        }
        return -1;
    }
}
