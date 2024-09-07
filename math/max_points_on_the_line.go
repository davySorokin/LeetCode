func maxPoints(points [][]int) int {
    // If there is only 1 point or none, the maximum number of points on a line is the number of points.
    if len(points) == 0 {
        return 0
    }
    if len(points) == 1 {
        return 1
    }
    
    // This function will calculate the greatest common divisor (GCD) of two integers.
    // This is used to reduce the slope's numerator and denominator to their simplest form.
    gcd := func(a, b int) int {
        for b != 0 {
            a, b = b, a % b
        }
        return a
    }
    
    // Initialize the maximum number of points on a single line.
    maxPointsOnLine := 0
    
    // Iterate through each point and treat it as the "base" point.
    for i := 0; i < len(points); i++ {
        // Create a map to store the slopes between the current base point and all other points.
        slopeMap := make(map[string]int)
        samePointCount := 1  // Count points that are exactly the same as the base point.
        currentMax := 0      // Tracks the maximum number of points for the current base point.
        
        // Compare the current base point with all the remaining points.
        for j := i + 1; j < len(points); j++ {
            dx := points[j][0] - points[i][0]
            dy := points[j][1] - points[i][1]
            
            // Check if the points are the same.
            if dx == 0 && dy == 0 {
                samePointCount++
                continue
            }
            
            // Reduce the slope to its simplest form by dividing both dx and dy by their GCD.
            d := gcd(dx, dy)
            dx /= d
            dy /= d
            
            // Ensure the slope is always stored consistently by ensuring the denominator is positive.
            // This avoids issues with (-1/2) and (1/-2) being treated as different slopes.
            if dx < 0 {
                dx = -dx
                dy = -dy
            }
            
            // Create a key for the slope in the format "dx/dy".
            slope := fmt.Sprintf("%d/%d", dx, dy)
            slopeMap[slope]++
            
            // Track the current maximum number of points with the same slope.
            if slopeMap[slope] > currentMax {
                currentMax = slopeMap[slope]
            }
        }
        
        // Update the global maximum points on a line, considering the points that are the same as the base point.
        if currentMax + samePointCount > maxPointsOnLine {
            maxPointsOnLine = currentMax + samePointCount
        }
    }
    
    // Return the maximum number of points on a line.
    return maxPointsOnLine
}
