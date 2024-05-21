package main

import "fmt"

func singleNumber(nums []int) int {
    result := 0
    for i := 0; i < 32; i++ {
        sum := 0
        for _, num := range nums {
            sum += (num >> i) & 1
        }
        result |= (sum % 3) << i
    }
    // Interpreting the result as a signed 32-bit integer
    if result >= 1<<31 {  // If the result is negative in two's complement
        result -= 1 << 32
    }
    return result
}
