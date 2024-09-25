package main

// main func wasn't necessary.

func circularPermutation(n int, start int) []int {
    // Generate the Gray code sequence for n bits
    size := 1 << n // This is 2^n
    grayCode := make([]int, size)
    
    for i := 0; i < size; i++ {
        grayCode[i] = i ^ (i >> 1) // Generate Gray code for i
    }
    
    // Find the index of the 'start' in the gray code sequence
    startIndex := 0
    for i := 0; i < size; i++ {
        if grayCode[i] == start {
            startIndex = i
            break
        }
    }
    
    // Rotate the sequence to start with the given 'start' element
    result := make([]int, size)
    for i := 0; i < size; i++ {
        result[i] = grayCode[(startIndex + i) % size]
    }
    
    return result
}
