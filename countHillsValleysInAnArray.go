func countHillValley(nums []int) int {
    // Remove consecutive duplicates
    newNums := []int{nums[0]}
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            newNums = append(newNums, nums[i])
        }
    }

    // Count hills and valleys
    count := 0
    for i := 1; i < len(newNums)-1; i++ {
        if (newNums[i] > newNums[i-1] && newNums[i] > newNums[i+1]) || 
           (newNums[i] < newNums[i-1] && newNums[i] < newNums[i+1]) {
            count++
        }
    }

    return count
}
