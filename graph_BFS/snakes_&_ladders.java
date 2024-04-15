class Solution {
    public int snakesAndLadders(int[][] board) {
        int n = board.length;
        boolean[] visited = new boolean[n * n + 1];
        Queue<Integer> queue = new LinkedList<>();
        
        queue.offer(1);
        visited[1] = true;
        int moves = 0;

        while (!queue.isEmpty()) {
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                int curr = queue.poll();
                if (curr == n * n) return moves; // Reached the end
                for (int j = 1; j <= 6 && curr + j <= n * n; j++) { // Roll the die
                    int next = curr + j;
                    int value = getBoardValue(board, next); // Get board value
                    if (value > 0) next = value; // Ladder or snake
                    if (!visited[next]) {
                        visited[next] = true;
                        queue.offer(next);
                    }
                }
            }
            moves++; // Increment moves after each level
        }
        
        return -1; // Impossible to reach the end
    }
    
    private int getBoardValue(int[][] board, int number) {
        int n = board.length;
        int r = (number - 1) / n;
        int x = n - 1 - r;
        int y = r % 2 == 0 ? (number - 1) % n : n - 1 - (number - 1) % n;
        return board[x][y];
    }
}
