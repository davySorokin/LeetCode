package main

import "fmt"

// longestValidParentheses returns the length of the longest valid (well-formed) parentheses substring.
func longestValidParentheses(s string) int {
    maxLen := 0
    stack := []int{-1}  // Start with -1 to handle the base for valid substring length calculation

    for i, char := range s {
        if char == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1] // Pop the last element
            if len(stack) == 0 {
                stack = append(stack, i) // Push current index as base if stack is empty
            } else {
                // Calculate the length if the stack is not empty
                maxLen = max(maxLen, i-stack[len(stack)-1])
            }
        }
    }
    return maxLen
}

// max returns the maximum of two integers.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
