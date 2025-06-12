func findErrorNums(nums []int) []int {
    n := len(nums)
    numMap := make(map[int]bool)
    duplicate := -1

    for _, num := range nums {
        if numMap[num] {
            duplicate = num
        }
        numMap[num] = true
    }

    for i := 1; i <= n; i++ {
        if !numMap[i] {
            return []int{duplicate, i}
        }
    }

    return nil
}
