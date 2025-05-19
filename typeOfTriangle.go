import "sort"

func triangleType(nums []int) string {
    sort.Ints(nums)
    a, b, c := nums[0], nums[1], nums[2]

    // Check triangle inequality
    if a + b <= c {
        return "none"
    }

    // Check for equilateral triangle
    if a == b && b == c {
        return "equilateral"
    }

    // Check for isosceles triangle
    if a == b || b == c || a == c {
        return "isosceles"
    }

    // All sides different
    return "scalene"
}
