package main

import (
    "fmt"
    "strings"
)

func addBinary(a string, b string) string {
    var result strings.Builder
    i, j := len(a)-1, len(b)-1
    carry := 0

    for i >= 0 || j >= 0 || carry != 0 {
        sum := carry
        if i >= 0 {
            sum += int(a[i] - '0') // Convert char to int ('0' or '1')
            i--
        }
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }
        carry = sum / 2        // Determine the new carry
        result.WriteByte('0' + byte(sum%2)) // Calculate the digit to be added to result
    }

    // The digits are added in reverse order, so we need to reverse the string before returning
    res := result.String()
    return reverseString(res)
}

// Helper function to reverse a string
func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
