package main

func findSmallestInteger(nums []int, value int) int {
    freq := make(map[int]int)
    for _, num := range nums {
        rem := ((num % value) + value) % value
        freq[rem]++
    }

    mex := 0
    for {
        rem := mex % value
        if freq[rem] > 0 {
            freq[rem]--
            mex++
        } else {
            return mex
        }
    }
}
