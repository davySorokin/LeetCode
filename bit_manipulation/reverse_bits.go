package main

import (
    "fmt"
)

func reverseBits(num uint32) uint32 {
    var result uint32
    for i := 0; i < 32; i++ {
        result <<= 1                // Shift result to the left to make room for the next bit
        result |= num & 1           // Add the least significant bit of num to result
        num >>= 1                   // Shift num to the right to process the next bit
    }
    return result
}
