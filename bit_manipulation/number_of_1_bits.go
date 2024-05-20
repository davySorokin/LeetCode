package main

import (
    "fmt"
)

func hammingWeight(n int) int {
    num := uint32(n)  // Convert to uint32 to handle bit operations correctly
    count := 0
    for num != 0 {
        num &= num - 1  // Drop the lowest set bit
        count++
    }
    return count
}
