func canMakeSquare(grid [][]byte) bool {

    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {

            countB, countW := 0, 0
            for x := i; x < i+2; x++ {
                for y := j; y < j+2; y++ {
                    if grid[x][y] == 'B' {
                        countB++
                    } else {
                        countW++
                    }
                }
            }
          
            if countB >= 3 || countW >= 3 {
                return true
            }
        }
    }
  
    return false
}
