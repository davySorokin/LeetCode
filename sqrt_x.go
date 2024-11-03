func mySqrt(x int) int {
	if x < 2 {
		return x
	}

  // left, right := 2, x / 2
	left, right := 1, x / 2  // Set left to 1 to handle small values like x = 2
	var result int

	for left <= right {
		mid := left + (right - left) / 2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
