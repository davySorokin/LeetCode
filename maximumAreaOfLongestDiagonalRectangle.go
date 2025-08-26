import (
    "math"
)

func areaOfMaxDiagonal(dimensions [][]int) int {
    maxDiag := 0.0
    maxArea := 0

    for _, dim := range dimensions {
        length := dim[0]
        width := dim[1]
        diag := math.Sqrt(float64(length*length + width*width))
        area := length * width

        if diag > maxDiag {
            maxDiag = diag
            maxArea = area
        } else if diag == maxDiag {
            if area > maxArea {
                maxArea = area
            }
        }
    }

    return maxArea
}

