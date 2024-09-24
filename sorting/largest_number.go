import (
    "sort"
    "strconv"
    "strings"
)

func largestNumber(nums []int) string {
    // Convert integers to strings
    strNums := make([]string, len(nums))
    for i, num := range nums {
        strNums[i] = strconv.Itoa(num)
    }

    // Custom sort the string numbers
    sort.Slice(strNums, func(i, j int) bool {
        // Compare by concatenating in both possible orders
        return strNums[i]+strNums[j] > strNums[j]+strNums[i]
    })

    // Edge case: if the largest number is "0", return "0"
    if strNums[0] == "0" {
        return "0"
    }

    // Join the sorted strings to form the largest number
    return strings.Join(strNums, "")
}
