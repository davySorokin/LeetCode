package main

import "fmt"

func longestMountain(arr []int) int {
    n := len(arr)
    if n < 3 {
        return 0
    }

    maxLen := 0
    i := 1

    for i < n-1 {
        // Check if the current element is a peak
        if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
            left := i
            right := i

            // Expand to the left
            for left > 0 && arr[left-1] < arr[left] {
                left--
            }

            // Expand to the right
            for right < n-1 && arr[right] > arr[right+1] {
                right++
            }

            // Calculate the length of the mountain
            currentLen := right - left + 1
            if currentLen > maxLen {
                maxLen = currentLen
            }

            // Move the index to the end of this mountain
            i = right
        } else {
            i++
        }
    }

    return maxLen
}
