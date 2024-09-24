func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    if indexDiff <= 0 || valueDiff < 0 {
        return false
    }

    bucketSize := valueDiff + 1
    buckets := make(map[int]int)

    for i, num := range nums {
        bucketID := getBucketID(num, bucketSize)

        // Check the current bucket
        if _, exists := buckets[bucketID]; exists {
            return true
        }

        // Check the neighboring buckets
        if neighbor, exists := buckets[bucketID-1]; exists && abs(num-neighbor) <= valueDiff {
            return true
        }
        if neighbor, exists := buckets[bucketID+1]; exists && abs(num-neighbor) <= valueDiff {
            return true
        }

        // Add the current number to its bucket
        buckets[bucketID] = num

        // Remove the element that is out of the `indexDiff` range
        if i >= indexDiff {
            delete(buckets, getBucketID(nums[i-indexDiff], bucketSize))
        }
    }

    return false
}

func getBucketID(num int, bucketSize int) int {
    if num < 0 {
        return (num+1)/bucketSize - 1
    }
    return num / bucketSize
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
