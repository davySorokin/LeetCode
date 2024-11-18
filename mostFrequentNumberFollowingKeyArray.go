func mostFrequent(nums []int, key int) int {
    
  count := make(map[int]int)
    
    // Traverse through the nums array
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == key {
            count[nums[i+1]]++
        }
    }
    
    // Find target with the maximum count
    maxCount := 0
    result := 0
    for target, freq := range count {
        if freq > maxCount {
            maxCount = freq
            result = target
        }
    }
    
    return result
}
