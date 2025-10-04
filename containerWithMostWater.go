func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0

    for left < right {
        h := min(height[left], height[right])
        w := right - left
        area := h * w
        if area > maxArea {
            maxArea = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
