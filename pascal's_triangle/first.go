package main

import "fmt"

func generate(numRows int) [][]int {
    triangle := make([][]int, numRows) // Initialize a 2D slice

    for i := 0; i < numRows; i++ {
        // Each row has `i + 1` elements
        row := make([]int, i+1)
        row[0], row[i] = 1, 1 // First and last elements are always 1

        // Fill in the middle elements
        for j := 1; j < i; j++ {
            row[j] = triangle[i-1][j-1] + triangle[i-1][j]
        }

        triangle[i] = row
    }

    return triangle
}

func main() {
    // Example 1
    numRows := 5
    fmt.Println(generate(numRows))

    // Example 2
    numRows = 1
    fmt.Println(generate(numRows))
}
