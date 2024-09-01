func numberOfSubarrays(nums []int, k int) int {
    count := 0
    oddCount := 0
    prefixCount := make(map[int]int)
    prefixCount[0] = 1  // There is one way to have a sum of 0

    for _, num := range nums {
        if num%2 == 1 { // num is odd
            oddCount++
        }
        if oddCount >= k {
            count += prefixCount[oddCount - k]
        }
        prefixCount[oddCount]++
    }
    
    return count
}
