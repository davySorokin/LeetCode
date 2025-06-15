package main

import "sort"

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    i, j, happy := 0, 0, 0
    for i < len(g) && j < len(s) {
        if s[j] >= g[i] {     // cookie j is big enough for child i
            happy++
            i++
            j++
        } else {              // cookie j too small, try a bigger one
            j++
        }
    }
    return happy
}
