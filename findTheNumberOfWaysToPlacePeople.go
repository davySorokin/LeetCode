package main

func numberOfPairs(points [][]int) int {
	n := len(points)
	ans := 0
	for i := 0; i < n; i++ {
		ax, ay := points[i][0], points[i][1]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			bx, by := points[j][0], points[j][1]
			// A is upper-left of B (lines allowed)
			if ax <= bx && ay >= by && (ax != bx || ay != by) {
				blocked := false
				minx, maxx := ax, bx
				miny, maxy := by, ay
				for k := 0; k < n; k++ {
					if k == i || k == j {
						continue
					}
					xk, yk := points[k][0], points[k][1]
					if xk >= minx && xk <= maxx && yk >= miny && yk <= maxy {
						blocked = true
						break
					}
				}
				if !blocked {
					ans++
				}
			}
		}
	}
	return ans
}
