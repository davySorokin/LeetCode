package main

import (
	"fmt"
)

func nthPersonGetsNthSeat(n int) float64 {
    if n == 1 {
        return 1.0
    }
    return 0.5
}

// test cases - not used in main program.
func main() {
    // Test cases
    fmt.Println(nthPersonGetsNthSeat(1)) // Expected: 1.00000
    fmt.Println(nthPersonGetsNthSeat(2)) // Expected: 0.50000
    fmt.Println(nthPersonGetsNthSeat(3)) // Expected: 0.50000
    fmt.Println(nthPersonGetsNthSeat(10)) // Expected: 0.50000
}
