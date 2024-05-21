package main

import "fmt"

func singleNumber(nums []int) int {
    var result int
    for i := 0; i < 32; i++ {
        sum := 0
        for _, num := range nums {
            sum += (num >> i) & 1
        }
        if sum % 3 != 0 {
            result |= (1 << i)
        }
    }
    return result
}
