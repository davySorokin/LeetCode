package main

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
    shifts := 0
    // Shift left and right to the right until they are the same
    for left != right {
        left >>= 1
        right >>= 1
        shifts++
    }
    // Shift back to the original position
    return left << shifts
}
