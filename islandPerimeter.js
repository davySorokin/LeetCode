/**
 * @param {number[][]} grid
 * @return {number}
 */
var islandPerimeter = function(grid) {
    let perimeter = 0;
    let rows = grid.length;
    let cols = grid[0].length;
    
    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (grid[r][c] === 1) {
                perimeter += 4;
                
                // Check adjacent cells and subtract for each shared edge
                if (r > 0 && grid[r - 1][c] === 1) perimeter -= 2; // Top
                if (c > 0 && grid[r][c - 1] === 1) perimeter -= 2; // Left
            }
        }
    }
    
    return perimeter;
};
