func maximumGap(nums []int) int {
    if len(nums) < 2 {
        return 0
    }

    // Find the minimum and maximum element in the array
    minNum, maxNum := nums[0], nums[0]
    for _, num := range nums {
        if num < minNum {
            minNum = num
        }
        if num > maxNum {
            maxNum = num
        }
    }

    // Handle edge case where all numbers are the same
    if minNum == maxNum {
        return 0
    }

    n := len(nums)
    // Calculate the minimum possible gap using bucket size
    bucketSize := max(1, (maxNum-minNum)/(n-1))
    bucketCount := (maxNum-minNum)/bucketSize + 1

    // Initialise buckets to store min and max values for each bucket
    minBuckets := make([]int, bucketCount)
    maxBuckets := make([]int, bucketCount)
    for i := range minBuckets {
        minBuckets[i] = -1 // Sentinel value for empty bucket
        maxBuckets[i] = -1
    }

    // Place each number in its corresponding bucket
    for _, num := range nums {
        bucketIndex := (num - minNum) / bucketSize
        if minBuckets[bucketIndex] == -1 {
            minBuckets[bucketIndex] = num
            maxBuckets[bucketIndex] = num
        } else {
            minBuckets[bucketIndex] = min(minBuckets[bucketIndex], num)
            maxBuckets[bucketIndex] = max(maxBuckets[bucketIndex], num)
        }
    }

    // Calculate the maximum gap by comparing adjacent non-empty buckets
    maxGap := 0
    previousMax := minNum
    for i := 0; i < bucketCount; i++ {
        if minBuckets[i] == -1 {
            continue // Skip empty buckets
        }
        maxGap = max(maxGap, minBuckets[i]-previousMax)
        previousMax = maxBuckets[i]
    }

    return maxGap
}

// Helper functions for min and max
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
