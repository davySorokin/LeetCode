import (
    "math"
)

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    // Use a map to simulate a sliding window of the last indexDiff elements
    window := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        // For every element nums[i], we search the window for nearby elements
        for _, v := range window {
            if math.Abs(float64(nums[i]-v)) <= float64(valueDiff) {
                return true
            }
        }

        // Add the current element to the window
        window[i] = nums[i]

        // Remove the oldest element that is out of the sliding window of size indexDiff
        if len(window) > indexDiff {
            delete(window, i-indexDiff)
        }
    }

    return false
}
