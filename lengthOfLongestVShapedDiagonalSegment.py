from typing import List

class Solution:
    def lenOfVDiagonal(self, grid: List[List[int]]) -> int:
        n, m = len(grid), len(grid[0])
        dirs = [(1,1), (1,-1), (-1,-1), (-1,1)]  # SE, SW, NW, NE
        clockwise = {0:1, 1:2, 2:3, 3:0}         # mapping of index
        
        def expected(step: int) -> int:
            if step == 0: return 1
            return 2 if step % 2 == 1 else 0
        
        best = 0
        
        for i in range(n):
            for j in range(m):
                if grid[i][j] != 1:
                    continue
                for d_idx, (dx,dy) in enumerate(dirs):
                    # simulate without turn and with turn
                    for turn_taken in [False]:
                        x, y = i, j
                        step = 0
                        cur_dir = d_idx
                        while 0 <= x < n and 0 <= y < m and grid[x][y] == expected(step):
                            best = max(best, step+1)
                            # decide direction
                            nx, ny = x + dirs[cur_dir][0], y + dirs[cur_dir][1]
                            # try to take turn exactly at this point if not taken
                            if not turn_taken:
                                # branch with turn
                                tx, ty = x + dirs[clockwise[cur_dir]][0], y + dirs[clockwise[cur_dir]][1]
                                t_step = step+1
                                t_dir = clockwise[cur_dir]
                                while 0 <= tx < n and 0 <= ty < m and grid[tx][ty] == expected(t_step):
                                    best = max(best, t_step+1)
                                    tx += dirs[t_dir][0]
                                    ty += dirs[t_dir][1]
                                    t_step += 1
                            # move forward
                            x, y = nx, ny
                            step += 1
        return best
