import (
    "strconv"
)

func getPermutation(n int, k int) string {
    factorial := 1
    nums := []int{}
    
    // prepare factorial and number list
    for i := 1; i <= n; i++ {
        factorial *= i
        nums = append(nums, i)
    }

    k-- // convert to 0-based index
    res := ""
    
    for i := n; i >= 1; i-- {
        factorial /= i
        index := k / factorial
        res += strconv.Itoa(nums[index])
        nums = append(nums[:index], nums[index+1:]...)
        k %= factorial
    }

    return res
}
