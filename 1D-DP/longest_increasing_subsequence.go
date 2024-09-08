import "sort"

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // This slice will keep track of the current increasing subsequence
    lis := []int{}

    for _, num := range nums {
        // Find the position where this number can fit in the current subsequence
        pos := sort.SearchInts(lis, num)

        if pos < len(lis) {
            // If the number can replace an element, we replace it
            lis[pos] = num
        } else {
            // If it's larger than all elements, we append it
            lis = append(lis, num)
        }
    }

    return len(lis)
}
