package main

import "fmt"

func getRow(rowIndex int) []int {
    row := make([]int, rowIndex+1)
    row[0] = 1 // First element is always 1

    for i := 1; i <= rowIndex; i++ {
        // Update the row from the end to the start to avoid overwriting values prematurely
        for j := i; j > 0; j-- {
            row[j] += row[j-1]
        }
    }

    return row
}

func main() {
    rowIndex := 3
    fmt.Println(getRow(rowIndex)) // Output: [1, 3, 3, 1]

    rowIndex = 0
    fmt.Println(getRow(rowIndex)) // Output: [1]

    rowIndex = 1
    fmt.Println(getRow(rowIndex)) // Output: [1, 1]
}
